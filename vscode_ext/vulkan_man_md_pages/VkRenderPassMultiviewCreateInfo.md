# VkRenderPassMultiviewCreateInfo(3) Manual Page

## Name

VkRenderPassMultiviewCreateInfo - Structure containing multiview
information for all subpasses



## <a href="#_c_specification" class="anchor"></a>C Specification

If the [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html)::`pNext`
chain includes a `VkRenderPassMultiviewCreateInfo` structure, then that
structure includes an array of view masks, view offsets, and correlation
masks for the render pass.

The `VkRenderPassMultiviewCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkRenderPassMultiviewCreateInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           subpassCount;
    const uint32_t*    pViewMasks;
    uint32_t           dependencyCount;
    const int32_t*     pViewOffsets;
    uint32_t           correlationMaskCount;
    const uint32_t*    pCorrelationMasks;
} VkRenderPassMultiviewCreateInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_multiview
typedef VkRenderPassMultiviewCreateInfo VkRenderPassMultiviewCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `subpassCount` is zero or the number of subpasses in the render pass.

- `pViewMasks` is a pointer to an array of `subpassCount` view masks,
  where each mask is a bitfield of view indices describing which views
  rendering is broadcast to in each subpass, when multiview is enabled.
  If `subpassCount` is zero, each view mask is treated as zero.

- `dependencyCount` is zero or the number of dependencies in the render
  pass.

- `pViewOffsets` is a pointer to an array of `dependencyCount` view
  offsets, one for each dependency. If `dependencyCount` is zero, each
  dependency’s view offset is treated as zero. Each view offset controls
  which views in the source subpass the views in the destination subpass
  depend on.

- `correlationMaskCount` is zero or the number of correlation masks.

- `pCorrelationMasks` is a pointer to an array of `correlationMaskCount`
  view masks indicating sets of views that **may** be more efficient to
  render concurrently.

## <a href="#_description" class="anchor"></a>Description

When a subpass uses a non-zero view mask, *multiview* functionality is
considered to be enabled. Multiview is all-or-nothing for a render
pass - that is, either all subpasses **must** have a non-zero view mask
(though some subpasses **may** have only one view) or all **must** be
zero. Multiview causes all drawing and clear commands in the subpass to
behave as if they were broadcast to each view, where a view is
represented by one layer of the framebuffer attachments. All draws and
clears are broadcast to each *view index* whose bit is set in the view
mask. The view index is provided in the `ViewIndex` shader input
variable, and color, depth/stencil, and input attachments all read/write
the layer of the framebuffer corresponding to the view index.

If the view mask is zero for all subpasses, multiview is considered to
be disabled and all drawing commands execute normally, without this
additional broadcasting.

Some implementations **may** not support multiview in conjunction with
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview-mesh"
target="_blank" rel="noopener">mesh shaders</a>, <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview-gs"
target="_blank" rel="noopener">geometry shaders</a> or <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview-tess"
target="_blank" rel="noopener">tessellation shaders</a>.

When multiview is enabled, the `VK_DEPENDENCY_VIEW_LOCAL_BIT` bit in a
dependency **can** be used to express a view-local dependency, meaning
that each view in the destination subpass depends on a single view in
the source subpass. Unlike pipeline barriers, a subpass dependency
**can** potentially have a different view mask in the source subpass and
the destination subpass. If the dependency is view-local, then each view
(dstView) in the destination subpass depends on the view dstView +
`pViewOffsets`\[dependency\] in the source subpass. If there is not such
a view in the source subpass, then this dependency does not affect that
view in the destination subpass. If the dependency is not view-local,
then all views in the destination subpass depend on all views in the
source subpass, and the view offset is ignored. A non-zero view offset
is not allowed in a self-dependency.

The elements of `pCorrelationMasks` are a set of masks of views
indicating that views in the same mask **may** exhibit spatial coherency
between the views, making it more efficient to render them concurrently.
Correlation masks **must** not have a functional effect on the results
of the multiview rendering.

When multiview is enabled, at the beginning of each subpass all
non-render pass state is undefined. In particular, each time
[vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass.html) or
[vkCmdNextSubpass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass.html) is called the graphics
pipeline **must** be bound, any relevant descriptor sets or vertex/index
buffers **must** be bound, and any relevant dynamic state or push
constants **must** be set before they are used.

A multiview subpass **can** declare that its shaders will write per-view
attributes for all views in a single invocation, by setting the
`VK_SUBPASS_DESCRIPTION_PER_VIEW_ATTRIBUTES_BIT_NVX` bit in the subpass
description. The only supported per-view attributes are position and
viewport mask, and per-view position and viewport masks are written to
output array variables decorated with `PositionPerViewNV` and
`ViewportMaskPerViewNV`, respectively. If
[`VK_NV_viewport_array2`](VK_NV_viewport_array2.html) is not supported
and enabled, `ViewportMaskPerViewNV` **must** not be used. Values
written to elements of `PositionPerViewNV` and `ViewportMaskPerViewNV`
**must** not depend on the `ViewIndex`. The shader **must** also write
to an output variable decorated with `Position`, and the value written
to `Position` **must** equal the value written to
`PositionPerViewNV`\[`ViewIndex`\]. Similarly, if
`ViewportMaskPerViewNV` is written to then the shader **must** also
write to an output variable decorated with `ViewportMaskNV`, and the
value written to `ViewportMaskNV` **must** equal the value written to
`ViewportMaskPerViewNV`\[`ViewIndex`\]. Implementations will either use
values taken from `Position` and `ViewportMaskNV` and invoke the shader
once for each view, or will use values taken from `PositionPerViewNV`
and `ViewportMaskPerViewNV` and invoke the shader fewer times. The
values written to `Position` and `ViewportMaskNV` **must** not depend on
the values written to `PositionPerViewNV` and `ViewportMaskPerViewNV`,
or vice versa (to allow compilers to eliminate the unused outputs). All
attributes that do not have `*PerViewNV` counterparts **must** not
depend on `ViewIndex`.

Per-view attributes are all-or-nothing for a subpass. That is, all
pipelines compiled against a subpass that includes the
`VK_SUBPASS_DESCRIPTION_PER_VIEW_ATTRIBUTES_BIT_NVX` bit **must** write
per-view attributes to the `*PerViewNV[]` shader outputs, in addition to
the non-per-view (e.g. `Position`) outputs. Pipelines compiled against a
subpass that does not include this bit **must** not include the
`*PerViewNV[]` outputs in their interfaces.

Valid Usage

- <a href="#VUID-VkRenderPassMultiviewCreateInfo-pCorrelationMasks-00841"
  id="VUID-VkRenderPassMultiviewCreateInfo-pCorrelationMasks-00841"></a>
  VUID-VkRenderPassMultiviewCreateInfo-pCorrelationMasks-00841  
  Each view index **must** not be set in more than one element of
  `pCorrelationMasks`

- <a href="#VUID-VkRenderPassMultiviewCreateInfo-multiview-06555"
  id="VUID-VkRenderPassMultiviewCreateInfo-multiview-06555"></a>
  VUID-VkRenderPassMultiviewCreateInfo-multiview-06555  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview"
  target="_blank" rel="noopener"><code>multiview</code></a> feature is
  not enabled, each element of `pViewMasks` **must** be `0`

- <a href="#VUID-VkRenderPassMultiviewCreateInfo-pViewMasks-06697"
  id="VUID-VkRenderPassMultiviewCreateInfo-pViewMasks-06697"></a>
  VUID-VkRenderPassMultiviewCreateInfo-pViewMasks-06697  
  The index of the most significant bit in each element of `pViewMasks`
  **must** be less than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxMultiviewViewCount"
  target="_blank" rel="noopener"><code>maxMultiviewViewCount</code></a>

Valid Usage (Implicit)

- <a href="#VUID-VkRenderPassMultiviewCreateInfo-sType-sType"
  id="VUID-VkRenderPassMultiviewCreateInfo-sType-sType"></a>
  VUID-VkRenderPassMultiviewCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO`

- <a href="#VUID-VkRenderPassMultiviewCreateInfo-pViewMasks-parameter"
  id="VUID-VkRenderPassMultiviewCreateInfo-pViewMasks-parameter"></a>
  VUID-VkRenderPassMultiviewCreateInfo-pViewMasks-parameter  
  If `subpassCount` is not `0`, `pViewMasks` **must** be a valid pointer
  to an array of `subpassCount` `uint32_t` values

- <a href="#VUID-VkRenderPassMultiviewCreateInfo-pViewOffsets-parameter"
  id="VUID-VkRenderPassMultiviewCreateInfo-pViewOffsets-parameter"></a>
  VUID-VkRenderPassMultiviewCreateInfo-pViewOffsets-parameter  
  If `dependencyCount` is not `0`, `pViewOffsets` **must** be a valid
  pointer to an array of `dependencyCount` `int32_t` values

- <a
  href="#VUID-VkRenderPassMultiviewCreateInfo-pCorrelationMasks-parameter"
  id="VUID-VkRenderPassMultiviewCreateInfo-pCorrelationMasks-parameter"></a>
  VUID-VkRenderPassMultiviewCreateInfo-pCorrelationMasks-parameter  
  If `correlationMaskCount` is not `0`, `pCorrelationMasks` **must** be
  a valid pointer to an array of `correlationMaskCount` `uint32_t`
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPassMultiviewCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
