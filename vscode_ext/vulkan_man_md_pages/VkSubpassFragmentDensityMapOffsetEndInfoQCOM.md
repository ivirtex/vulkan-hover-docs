# VkSubpassFragmentDensityMapOffsetEndInfoQCOM(3) Manual Page

## Name

VkSubpassFragmentDensityMapOffsetEndInfoQCOM - Structure specifying
fragment density map offset subpass end information



## <a href="#_c_specification" class="anchor"></a>C Specification

If the [VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassEndInfo.html)::`pNext` chain includes
a `VkSubpassFragmentDensityMapOffsetEndInfoQCOM` structure, then that
structure includes an array of fragment density map offsets per layer
for the render pass.

The `VkSubpassFragmentDensityMapOffsetEndInfoQCOM` structure is defined
as:

``` c
// Provided by VK_QCOM_fragment_density_map_offset
typedef struct VkSubpassFragmentDensityMapOffsetEndInfoQCOM {
    VkStructureType      sType;
    const void*          pNext;
    uint32_t             fragmentDensityOffsetCount;
    const VkOffset2D*    pFragmentDensityOffsets;
} VkSubpassFragmentDensityMapOffsetEndInfoQCOM;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `fragmentDensityOffsetCount` is the number of offsets being specified.

- `pFragmentDensityOffsets` is a pointer to an array of
  [VkOffset2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset2D.html) structs, each of which describes the
  offset per layer.

## <a href="#_description" class="anchor"></a>Description

The array elements are given per `layer` as defined by <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragmentdensitymap-fetch-density-value"
target="_blank" rel="noopener">Fetch Density Value</a>, where index =
layer. Each (x,y) offset is in framebuffer pixels and shifts the fetch
of the fragment density map by that amount. Offsets can be positive or
negative.

Offset values specified for any subpass that is not the last subpass in
the render pass are ignored. If the
[VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassEndInfo.html)::`pNext` chain for the last
subpass of a render pass does not include
`VkSubpassFragmentDensityMapOffsetEndInfoQCOM`, or if
`fragmentDensityOffsetCount` is zero, then the offset (0,0) is used for
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragmentdensitymap-fetch-density-value"
target="_blank" rel="noopener">Fetch Density Value</a>.

Valid Usage

- <a
  href="#VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-fragmentDensityMapOffsets-06503"
  id="VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-fragmentDensityMapOffsets-06503"></a>
  VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-fragmentDensityMapOffsets-06503  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fragmentDensityMapOffsets"
  target="_blank"
  rel="noopener"><code>fragmentDensityMapOffsets</code></a> feature is
  not enabled or fragment density map is not enabled in the render pass,
  `fragmentDensityOffsetCount` **must** equal `0`

- <a
  href="#VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-fragmentDensityMapAttachment-06504"
  id="VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-fragmentDensityMapAttachment-06504"></a>
  VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-fragmentDensityMapAttachment-06504  
  If `VkSubpassDescription`::`fragmentDensityMapAttachment` is not is
  not `VK_ATTACHMENT_UNUSED` and was not created with
  `VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_QCOM`,
  `fragmentDensityOffsetCount` **must** equal `0`

- <a
  href="#VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pDepthStencilAttachment-06505"
  id="VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pDepthStencilAttachment-06505"></a>
  VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pDepthStencilAttachment-06505  
  If `VkSubpassDescription`::`pDepthStencilAttachment` is not is not
  `VK_ATTACHMENT_UNUSED` and was not created with
  `VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_QCOM`,
  `fragmentDensityOffsetCount` **must** equal `0`

- <a
  href="#VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pInputAttachments-06506"
  id="VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pInputAttachments-06506"></a>
  VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pInputAttachments-06506  
  If any element of `VkSubpassDescription`::`pInputAttachments` is not
  is not `VK_ATTACHMENT_UNUSED` and was not created with
  `VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_QCOM`,
  `fragmentDensityOffsetCount` **must** equal `0`

- <a
  href="#VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pColorAttachments-06507"
  id="VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pColorAttachments-06507"></a>
  VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pColorAttachments-06507  
  If any element of `VkSubpassDescription`::`pColorAttachments` is not
  is not `VK_ATTACHMENT_UNUSED` and was not created with
  `VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_QCOM`,
  `fragmentDensityOffsetCount` **must** equal `0`

- <a
  href="#VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pResolveAttachments-06508"
  id="VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pResolveAttachments-06508"></a>
  VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pResolveAttachments-06508  
  If any element of `VkSubpassDescription`::`pResolveAttachments` is not
  is not `VK_ATTACHMENT_UNUSED` and was not created with
  `VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_QCOM`,
  `fragmentDensityOffsetCount` **must** equal `0`

- <a
  href="#VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pPreserveAttachments-06509"
  id="VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pPreserveAttachments-06509"></a>
  VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pPreserveAttachments-06509  
  If any element of `VkSubpassDescription`::`pPreserveAttachments` is
  not is not `VK_ATTACHMENT_UNUSED` and was not created with
  `VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_QCOM`,
  `fragmentDensityOffsetCount` **must** equal `0`

- <a
  href="#VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-fragmentDensityOffsetCount-06510"
  id="VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-fragmentDensityOffsetCount-06510"></a>
  VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-fragmentDensityOffsetCount-06510  
  If `fragmentDensityOffsetCount` is not `0` and multiview is enabled
  for the render pass, `fragmentDensityOffsetCount` **must** equal the
  `layerCount` that was specified in creating the fragment density map
  attachment view

- <a
  href="#VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-fragmentDensityOffsetCount-06511"
  id="VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-fragmentDensityOffsetCount-06511"></a>
  VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-fragmentDensityOffsetCount-06511  
  If `fragmentDensityOffsetCount` is not `0` and multiview is not
  enabled for the render pass, `fragmentDensityOffsetCount` **must**
  equal `1`

- <a href="#VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-x-06512"
  id="VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-x-06512"></a>
  VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-x-06512  
  The `x` component of each element of `pFragmentDensityOffsets`
  **must** be an integer multiple of
  `fragmentDensityOffsetGranularity.width`

- <a href="#VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-y-06513"
  id="VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-y-06513"></a>
  VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-y-06513  
  The `y` component of each element of `pFragmentDensityOffsets`
  **must** be an integer multiple of
  `fragmentDensityOffsetGranularity.height`

Valid Usage (Implicit)

- <a href="#VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-sType-sType"
  id="VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-sType-sType"></a>
  VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SUBPASS_FRAGMENT_DENSITY_MAP_OFFSET_END_INFO_QCOM`

- <a
  href="#VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pFragmentDensityOffsets-parameter"
  id="VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pFragmentDensityOffsets-parameter"></a>
  VUID-VkSubpassFragmentDensityMapOffsetEndInfoQCOM-pFragmentDensityOffsets-parameter  
  If `fragmentDensityOffsetCount` is not `0`, `pFragmentDensityOffsets`
  **must** be a valid pointer to an array of
  `fragmentDensityOffsetCount` [VkOffset2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset2D.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_fragment_density_map_offset](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_fragment_density_map_offset.html),
[VkOffset2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubpassFragmentDensityMapOffsetEndInfoQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
