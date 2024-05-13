# VkFragmentShadingRateAttachmentInfoKHR(3) Manual Page

## Name

VkFragmentShadingRateAttachmentInfoKHR - Structure specifying a fragment
shading rate attachment for a subpass



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkFragmentShadingRateAttachmentInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_fragment_shading_rate
typedef struct VkFragmentShadingRateAttachmentInfoKHR {
    VkStructureType                  sType;
    const void*                      pNext;
    const VkAttachmentReference2*    pFragmentShadingRateAttachment;
    VkExtent2D                       shadingRateAttachmentTexelSize;
} VkFragmentShadingRateAttachmentInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pFragmentShadingRateAttachment` is `NULL` or a pointer to a
  [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html) structure
  defining the fragment shading rate attachment for this subpass.

- `shadingRateAttachmentTexelSize` specifies the size of the portion of
  the framebuffer corresponding to each texel in
  `pFragmentShadingRateAttachment`.

## <a href="#_description" class="anchor"></a>Description

If no shading rate attachment is specified, or if this structure is not
specified, the implementation behaves as if a valid shading rate
attachment was specified with all texels specifying a single pixel per
fragment.

Valid Usage

- <a
  href="#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04524"
  id="VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04524"></a>
  VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04524  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment`
  member is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** be
  equal to `VK_IMAGE_LAYOUT_GENERAL` or
  `VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR`

- <a
  href="#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04525"
  id="VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04525"></a>
  VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04525  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment`
  member is not `VK_ATTACHMENT_UNUSED`,
  `shadingRateAttachmentTexelSize.width` **must** be a power of two
  value

- <a
  href="#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04526"
  id="VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04526"></a>
  VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04526  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment`
  member is not `VK_ATTACHMENT_UNUSED`,
  `shadingRateAttachmentTexelSize.width` **must** be less than or equal
  to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxFragmentShadingRateAttachmentTexelSize"
  target="_blank"
  rel="noopener"><code>maxFragmentShadingRateAttachmentTexelSize.width</code></a>

- <a
  href="#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04527"
  id="VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04527"></a>
  VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04527  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment`
  member is not `VK_ATTACHMENT_UNUSED`,
  `shadingRateAttachmentTexelSize.width` **must** be greater than or
  equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-minFragmentShadingRateAttachmentTexelSize"
  target="_blank"
  rel="noopener"><code>minFragmentShadingRateAttachmentTexelSize.width</code></a>

- <a
  href="#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04528"
  id="VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04528"></a>
  VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04528  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment`
  member is not `VK_ATTACHMENT_UNUSED`,
  `shadingRateAttachmentTexelSize.height` **must** be a power of two
  value

- <a
  href="#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04529"
  id="VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04529"></a>
  VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04529  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment`
  member is not `VK_ATTACHMENT_UNUSED`,
  `shadingRateAttachmentTexelSize.height` **must** be less than or equal
  to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxFragmentShadingRateAttachmentTexelSize"
  target="_blank"
  rel="noopener"><code>maxFragmentShadingRateAttachmentTexelSize.height</code></a>

- <a
  href="#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04530"
  id="VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04530"></a>
  VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04530  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment`
  member is not `VK_ATTACHMENT_UNUSED`,
  `shadingRateAttachmentTexelSize.height` **must** be greater than or
  equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-minFragmentShadingRateAttachmentTexelSize"
  target="_blank"
  rel="noopener"><code>minFragmentShadingRateAttachmentTexelSize.height</code></a>

- <a
  href="#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04531"
  id="VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04531"></a>
  VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04531  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment`
  member is not `VK_ATTACHMENT_UNUSED`, the quotient of
  `shadingRateAttachmentTexelSize.width` and
  `shadingRateAttachmentTexelSize.height` **must** be less than or equal
  to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxFragmentShadingRateAttachmentTexelSizeAspectRatio"
  target="_blank"
  rel="noopener"><code>maxFragmentShadingRateAttachmentTexelSizeAspectRatio</code></a>

- <a
  href="#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04532"
  id="VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04532"></a>
  VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04532  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment`
  member is not `VK_ATTACHMENT_UNUSED`, the quotient of
  `shadingRateAttachmentTexelSize.height` and
  `shadingRateAttachmentTexelSize.width` **must** be less than or equal
  to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxFragmentShadingRateAttachmentTexelSizeAspectRatio"
  target="_blank"
  rel="noopener"><code>maxFragmentShadingRateAttachmentTexelSizeAspectRatio</code></a>

Valid Usage (Implicit)

- <a href="#VUID-VkFragmentShadingRateAttachmentInfoKHR-sType-sType"
  id="VUID-VkFragmentShadingRateAttachmentInfoKHR-sType-sType"></a>
  VUID-VkFragmentShadingRateAttachmentInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_FRAGMENT_SHADING_RATE_ATTACHMENT_INFO_KHR`

- <a
  href="#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-parameter"
  id="VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-parameter"></a>
  VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-parameter  
  If `pFragmentShadingRateAttachment` is not `NULL`,
  `pFragmentShadingRateAttachment` **must** be a valid pointer to a
  valid [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_fragment_shading_rate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_fragment_shading_rate.html),
[VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFragmentShadingRateAttachmentInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
