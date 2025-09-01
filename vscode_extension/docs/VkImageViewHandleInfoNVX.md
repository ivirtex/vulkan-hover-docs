# VkImageViewHandleInfoNVX(3) Manual Page

## Name

VkImageViewHandleInfoNVX - Structure specifying the image view for handle queries



## [](#_c_specification)C Specification

The `VkImageViewHandleInfoNVX` structure is defined as:

```c++
// Provided by VK_NVX_image_view_handle
typedef struct VkImageViewHandleInfoNVX {
    VkStructureType     sType;
    const void*         pNext;
    VkImageView         imageView;
    VkDescriptorType    descriptorType;
    VkSampler           sampler;
} VkImageViewHandleInfoNVX;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `imageView` is the image view to query.
- `descriptorType` is the type of descriptor for which to query a handle.
- `sampler` is the sampler to combine with the image view when generating the handle.

## [](#_description)Description

Valid Usage

- [](#VUID-VkImageViewHandleInfoNVX-descriptorType-02654)VUID-VkImageViewHandleInfoNVX-descriptorType-02654  
  `descriptorType` **must** be `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`
- [](#VUID-VkImageViewHandleInfoNVX-sampler-02655)VUID-VkImageViewHandleInfoNVX-sampler-02655  
  `sampler` **must** be a valid [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) if `descriptorType` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`
- [](#VUID-VkImageViewHandleInfoNVX-imageView-02656)VUID-VkImageViewHandleInfoNVX-imageView-02656  
  If descriptorType is `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` or `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, the image that `imageView` was created from **must** have been created with the `VK_IMAGE_USAGE_SAMPLED_BIT` usage bit set
- [](#VUID-VkImageViewHandleInfoNVX-imageView-02657)VUID-VkImageViewHandleInfoNVX-imageView-02657  
  If descriptorType is `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, the image that `imageView` was created from **must** have been created with the `VK_IMAGE_USAGE_STORAGE_BIT` usage bit set

Valid Usage (Implicit)

- [](#VUID-VkImageViewHandleInfoNVX-sType-sType)VUID-VkImageViewHandleInfoNVX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_VIEW_HANDLE_INFO_NVX`
- [](#VUID-VkImageViewHandleInfoNVX-pNext-pNext)VUID-VkImageViewHandleInfoNVX-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkImageViewHandleInfoNVX-imageView-parameter)VUID-VkImageViewHandleInfoNVX-imageView-parameter  
  `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle
- [](#VUID-VkImageViewHandleInfoNVX-descriptorType-parameter)VUID-VkImageViewHandleInfoNVX-descriptorType-parameter  
  `descriptorType` **must** be a valid [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html) value
- [](#VUID-VkImageViewHandleInfoNVX-sampler-parameter)VUID-VkImageViewHandleInfoNVX-sampler-parameter  
  If `sampler` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `sampler` **must** be a valid [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) handle
- [](#VUID-VkImageViewHandleInfoNVX-commonparent)VUID-VkImageViewHandleInfoNVX-commonparent  
  Both of `imageView`, and `sampler` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_NVX\_image\_view\_handle](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NVX_image_view_handle.html), [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetImageViewHandle64NVX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageViewHandle64NVX.html), [vkGetImageViewHandleNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageViewHandleNVX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageViewHandleInfoNVX).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0