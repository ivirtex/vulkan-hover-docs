# VkImageViewCaptureDescriptorDataInfoEXT(3) Manual Page

## Name

VkImageViewCaptureDescriptorDataInfoEXT - Structure specifying an image view for descriptor capture



## [](#_c_specification)C Specification

Information about the image view to get descriptor buffer capture data for is passed in a `VkImageViewCaptureDescriptorDataInfoEXT` structure:

```c++
// Provided by VK_EXT_descriptor_buffer
typedef struct VkImageViewCaptureDescriptorDataInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkImageView        imageView;
} VkImageViewCaptureDescriptorDataInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `imageView` is the `VkImageView` handle of the image view to get opaque capture data for.

## [](#_description)Description

Valid Usage

- [](#VUID-VkImageViewCaptureDescriptorDataInfoEXT-imageView-08083)VUID-VkImageViewCaptureDescriptorDataInfoEXT-imageView-08083  
  `imageView` **must** have been created with `VK_IMAGE_VIEW_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` set in [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html)::`flags`

Valid Usage (Implicit)

- [](#VUID-VkImageViewCaptureDescriptorDataInfoEXT-sType-sType)VUID-VkImageViewCaptureDescriptorDataInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_VIEW_CAPTURE_DESCRIPTOR_DATA_INFO_EXT`
- [](#VUID-VkImageViewCaptureDescriptorDataInfoEXT-pNext-pNext)VUID-VkImageViewCaptureDescriptorDataInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkImageViewCaptureDescriptorDataInfoEXT-imageView-parameter)VUID-VkImageViewCaptureDescriptorDataInfoEXT-imageView-parameter  
  `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetImageViewOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageViewOpaqueCaptureDescriptorDataEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageViewCaptureDescriptorDataInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0