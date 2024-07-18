# VkRenderingFragmentShadingRateAttachmentInfoKHR(3) Manual Page

## Name

VkRenderingFragmentShadingRateAttachmentInfoKHR - Structure specifying
fragment shading rate attachment information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRenderingFragmentShadingRateAttachmentInfoKHR` structure is
defined as:

``` c
// Provided by VK_KHR_dynamic_rendering with VK_KHR_fragment_shading_rate
typedef struct VkRenderingFragmentShadingRateAttachmentInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkImageView        imageView;
    VkImageLayout      imageLayout;
    VkExtent2D         shadingRateAttachmentTexelSize;
} VkRenderingFragmentShadingRateAttachmentInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `imageView` is the image view that will be used as a fragment shading
  rate attachment.

- `imageLayout` is the layout that `imageView` will be in during
  rendering.

- `shadingRateAttachmentTexelSize` specifies the number of pixels
  corresponding to each texel in `imageView`.

## <a href="#_description" class="anchor"></a>Description

This structure can be included in the `pNext` chain of
[VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html) to define a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-attachment"
target="_blank" rel="noopener">fragment shading rate attachment</a>. If
`imageView` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), or if this
structure is not specified, the implementation behaves as if a valid
shading rate attachment was specified with all texels specifying a
single pixel per fragment.

Valid Usage

- <a
  href="#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06147"
  id="VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06147"></a>
  VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06147  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `layout`
  **must** be `VK_IMAGE_LAYOUT_GENERAL` or
  `VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR`

- <a
  href="#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06148"
  id="VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06148"></a>
  VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06148  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it
  **must** have been created with
  `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a
  href="#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06149"
  id="VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06149"></a>
  VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06149  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `shadingRateAttachmentTexelSize.width` **must** be a power of two
  value

- <a
  href="#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06150"
  id="VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06150"></a>
  VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06150  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `shadingRateAttachmentTexelSize.width` **must** be less than or equal
  to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxFragmentShadingRateAttachmentTexelSize"
  target="_blank"
  rel="noopener"><code>maxFragmentShadingRateAttachmentTexelSize.width</code></a>

- <a
  href="#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06151"
  id="VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06151"></a>
  VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06151  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `shadingRateAttachmentTexelSize.width` **must** be greater than or
  equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-minFragmentShadingRateAttachmentTexelSize"
  target="_blank"
  rel="noopener"><code>minFragmentShadingRateAttachmentTexelSize.width</code></a>

- <a
  href="#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06152"
  id="VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06152"></a>
  VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06152  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `shadingRateAttachmentTexelSize.height` **must** be a power of two
  value

- <a
  href="#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06153"
  id="VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06153"></a>
  VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06153  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `shadingRateAttachmentTexelSize.height` **must** be less than or equal
  to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxFragmentShadingRateAttachmentTexelSize"
  target="_blank"
  rel="noopener"><code>maxFragmentShadingRateAttachmentTexelSize.height</code></a>

- <a
  href="#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06154"
  id="VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06154"></a>
  VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06154  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `shadingRateAttachmentTexelSize.height` **must** be greater than or
  equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-minFragmentShadingRateAttachmentTexelSize"
  target="_blank"
  rel="noopener"><code>minFragmentShadingRateAttachmentTexelSize.height</code></a>

- <a
  href="#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06155"
  id="VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06155"></a>
  VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06155  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  quotient of `shadingRateAttachmentTexelSize.width` and
  `shadingRateAttachmentTexelSize.height` **must** be less than or equal
  to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxFragmentShadingRateAttachmentTexelSizeAspectRatio"
  target="_blank"
  rel="noopener"><code>maxFragmentShadingRateAttachmentTexelSizeAspectRatio</code></a>

- <a
  href="#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06156"
  id="VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06156"></a>
  VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06156  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  quotient of `shadingRateAttachmentTexelSize.height` and
  `shadingRateAttachmentTexelSize.width` **must** be less than or equal
  to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxFragmentShadingRateAttachmentTexelSizeAspectRatio"
  target="_blank"
  rel="noopener"><code>maxFragmentShadingRateAttachmentTexelSizeAspectRatio</code></a>

Valid Usage (Implicit)

- <a
  href="#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-sType-sType"
  id="VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-sType-sType"></a>
  VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RENDERING_FRAGMENT_SHADING_RATE_ATTACHMENT_INFO_KHR`

- <a
  href="#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-parameter"
  id="VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-parameter"></a>
  VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-parameter  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handle

- <a
  href="#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageLayout-parameter"
  id="VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageLayout-parameter"></a>
  VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageLayout-parameter  
  `imageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_dynamic_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering.html),
[VK_KHR_fragment_shading_rate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_fragment_shading_rate.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderingFragmentShadingRateAttachmentInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
