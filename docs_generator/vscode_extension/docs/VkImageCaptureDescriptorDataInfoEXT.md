# VkImageCaptureDescriptorDataInfoEXT(3) Manual Page

## Name

VkImageCaptureDescriptorDataInfoEXT - Structure specifying an image for descriptor capture



## [](#_c_specification)C Specification

Information about the image to get descriptor buffer capture data for is passed in a `VkImageCaptureDescriptorDataInfoEXT` structure:

```c++
// Provided by VK_EXT_descriptor_buffer
typedef struct VkImageCaptureDescriptorDataInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkImage            image;
} VkImageCaptureDescriptorDataInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `image` is the `VkImage` handle of the image to get opaque capture data for.

## [](#_description)Description

Valid Usage

- [](#VUID-VkImageCaptureDescriptorDataInfoEXT-image-08079)VUID-VkImageCaptureDescriptorDataInfoEXT-image-08079  
  `image` **must** have been created with `VK_IMAGE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` set in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`flags`

Valid Usage (Implicit)

- [](#VUID-VkImageCaptureDescriptorDataInfoEXT-sType-sType)VUID-VkImageCaptureDescriptorDataInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_CAPTURE_DESCRIPTOR_DATA_INFO_EXT`
- [](#VUID-VkImageCaptureDescriptorDataInfoEXT-pNext-pNext)VUID-VkImageCaptureDescriptorDataInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkImageCaptureDescriptorDataInfoEXT-image-parameter)VUID-VkImageCaptureDescriptorDataInfoEXT-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetImageOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageOpaqueCaptureDescriptorDataEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageCaptureDescriptorDataInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0