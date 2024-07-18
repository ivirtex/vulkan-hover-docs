# VkImageViewHandleInfoNVX(3) Manual Page

## Name

VkImageViewHandleInfoNVX - Structure specifying the image view for
handle queries



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageViewHandleInfoNVX` structure is defined as:

``` c
// Provided by VK_NVX_image_view_handle
typedef struct VkImageViewHandleInfoNVX {
    VkStructureType     sType;
    const void*         pNext;
    VkImageView         imageView;
    VkDescriptorType    descriptorType;
    VkSampler           sampler;
} VkImageViewHandleInfoNVX;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `imageView` is the image view to query.

- `descriptorType` is the type of descriptor for which to query a
  handle.

- `sampler` is the sampler to combine with the image view when
  generating the handle.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkImageViewHandleInfoNVX-descriptorType-02654"
  id="VUID-VkImageViewHandleInfoNVX-descriptorType-02654"></a>
  VUID-VkImageViewHandleInfoNVX-descriptorType-02654  
  `descriptorType` **must** be `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`,
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`

- <a href="#VUID-VkImageViewHandleInfoNVX-sampler-02655"
  id="VUID-VkImageViewHandleInfoNVX-sampler-02655"></a>
  VUID-VkImageViewHandleInfoNVX-sampler-02655  
  `sampler` **must** be a valid [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) if
  `descriptorType` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`

- <a href="#VUID-VkImageViewHandleInfoNVX-imageView-02656"
  id="VUID-VkImageViewHandleInfoNVX-imageView-02656"></a>
  VUID-VkImageViewHandleInfoNVX-imageView-02656  
  If descriptorType is `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` or
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, the image that
  `imageView` was created from **must** have been created with the
  `VK_IMAGE_USAGE_SAMPLED_BIT` usage bit set

- <a href="#VUID-VkImageViewHandleInfoNVX-imageView-02657"
  id="VUID-VkImageViewHandleInfoNVX-imageView-02657"></a>
  VUID-VkImageViewHandleInfoNVX-imageView-02657  
  If descriptorType is `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, the image
  that `imageView` was created from **must** have been created with the
  `VK_IMAGE_USAGE_STORAGE_BIT` usage bit set

Valid Usage (Implicit)

- <a href="#VUID-VkImageViewHandleInfoNVX-sType-sType"
  id="VUID-VkImageViewHandleInfoNVX-sType-sType"></a>
  VUID-VkImageViewHandleInfoNVX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_VIEW_HANDLE_INFO_NVX`

- <a href="#VUID-VkImageViewHandleInfoNVX-pNext-pNext"
  id="VUID-VkImageViewHandleInfoNVX-pNext-pNext"></a>
  VUID-VkImageViewHandleInfoNVX-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkImageViewHandleInfoNVX-imageView-parameter"
  id="VUID-VkImageViewHandleInfoNVX-imageView-parameter"></a>
  VUID-VkImageViewHandleInfoNVX-imageView-parameter  
  `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handle

- <a href="#VUID-VkImageViewHandleInfoNVX-descriptorType-parameter"
  id="VUID-VkImageViewHandleInfoNVX-descriptorType-parameter"></a>
  VUID-VkImageViewHandleInfoNVX-descriptorType-parameter  
  `descriptorType` **must** be a valid
  [VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html) value

- <a href="#VUID-VkImageViewHandleInfoNVX-sampler-parameter"
  id="VUID-VkImageViewHandleInfoNVX-sampler-parameter"></a>
  VUID-VkImageViewHandleInfoNVX-sampler-parameter  
  If `sampler` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `sampler`
  **must** be a valid [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) handle

- <a href="#VUID-VkImageViewHandleInfoNVX-commonparent"
  id="VUID-VkImageViewHandleInfoNVX-commonparent"></a>
  VUID-VkImageViewHandleInfoNVX-commonparent  
  Both of `imageView`, and `sampler` that are valid handles of
  non-ignored parameters **must** have been created, allocated, or
  retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NVX_image_view_handle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NVX_image_view_handle.html),
[VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html),
[VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html), [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetImageViewHandleNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageViewHandleNVX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageViewHandleInfoNVX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
