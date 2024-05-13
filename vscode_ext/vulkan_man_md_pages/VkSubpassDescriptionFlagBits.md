# VkSubpassDescriptionFlagBits(3) Manual Page

## Name

VkSubpassDescriptionFlagBits - Bitmask specifying usage of a subpass



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkSubpassDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription.html)::`flags`, specifying
usage of the subpass, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkSubpassDescriptionFlagBits {
  // Provided by VK_NVX_multiview_per_view_attributes
    VK_SUBPASS_DESCRIPTION_PER_VIEW_ATTRIBUTES_BIT_NVX = 0x00000001,
  // Provided by VK_NVX_multiview_per_view_attributes
    VK_SUBPASS_DESCRIPTION_PER_VIEW_POSITION_X_ONLY_BIT_NVX = 0x00000002,
  // Provided by VK_QCOM_render_pass_shader_resolve
    VK_SUBPASS_DESCRIPTION_FRAGMENT_REGION_BIT_QCOM = 0x00000004,
  // Provided by VK_QCOM_render_pass_shader_resolve
    VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM = 0x00000008,
  // Provided by VK_EXT_rasterization_order_attachment_access
    VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_COLOR_ACCESS_BIT_EXT = 0x00000010,
  // Provided by VK_EXT_rasterization_order_attachment_access
    VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_DEPTH_ACCESS_BIT_EXT = 0x00000020,
  // Provided by VK_EXT_rasterization_order_attachment_access
    VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_STENCIL_ACCESS_BIT_EXT = 0x00000040,
  // Provided by VK_EXT_legacy_dithering
    VK_SUBPASS_DESCRIPTION_ENABLE_LEGACY_DITHERING_BIT_EXT = 0x00000080,
  // Provided by VK_ARM_rasterization_order_attachment_access
    VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_COLOR_ACCESS_BIT_ARM = VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_COLOR_ACCESS_BIT_EXT,
  // Provided by VK_ARM_rasterization_order_attachment_access
    VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_DEPTH_ACCESS_BIT_ARM = VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_DEPTH_ACCESS_BIT_EXT,
  // Provided by VK_ARM_rasterization_order_attachment_access
    VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_STENCIL_ACCESS_BIT_ARM = VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_STENCIL_ACCESS_BIT_EXT,
} VkSubpassDescriptionFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SUBPASS_DESCRIPTION_PER_VIEW_ATTRIBUTES_BIT_NVX` specifies that
  shaders compiled for this subpass write the attributes for all views
  in a single invocation of each <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader stage</a>. All
  pipelines compiled against a subpass that includes this bit **must**
  write per-view attributes to the `*PerViewNV[]` shader outputs, in
  addition to the non-per-view (e.g. `Position`) outputs.

- `VK_SUBPASS_DESCRIPTION_PER_VIEW_POSITION_X_ONLY_BIT_NVX` specifies
  that shaders compiled for this subpass use per-view positions which
  only differ in value in the x component. Per-view viewport mask
  **can** also be used.

- `VK_SUBPASS_DESCRIPTION_FRAGMENT_REGION_BIT_QCOM` specifies that the
  framebuffer region is the fragment region, that is, the minimum region
  dependencies are by pixel rather than by sample, such that any
  fragment shader invocation **can** access any sample associated with
  that fragment shader invocation.

- `VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM` specifies that the
  subpass performs shader resolve operations.

- `VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_COLOR_ACCESS_BIT_EXT`
  specifies that this subpass supports pipelines created with
  `VK_PIPELINE_COLOR_BLEND_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_ACCESS_BIT_EXT`.

- `VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_DEPTH_ACCESS_BIT_EXT`
  specifies that this subpass supports pipelines created with
  `VK_PIPELINE_DEPTH_STENCIL_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_DEPTH_ACCESS_BIT_EXT`.

- `VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_STENCIL_ACCESS_BIT_EXT`
  specifies that this subpass supports pipelines created with
  `VK_PIPELINE_DEPTH_STENCIL_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_STENCIL_ACCESS_BIT_EXT`.

- `VK_SUBPASS_DESCRIPTION_ENABLE_LEGACY_DITHERING_BIT_EXT` specifies
  that <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-legacy-dithering"
  target="_blank" rel="noopener">Legacy Dithering</a> is enabled for
  this subpass.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Shader resolve operations allow for custom resolve operations, but
overdrawing pixels <strong>may</strong> have a performance and/or power
cost. Furthermore, since the content of any depth stencil attachment or
color attachment is undefined at the beginning of a shader resolve
subpass, any depth testing, stencil testing, or blending operation which
sources these undefined values also has undefined result value.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkSubpassDescriptionFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubpassDescriptionFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
