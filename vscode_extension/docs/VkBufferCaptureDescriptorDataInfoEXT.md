# VkBufferCaptureDescriptorDataInfoEXT(3) Manual Page

## Name

VkBufferCaptureDescriptorDataInfoEXT - Structure specifying a buffer for descriptor capture



## [](#_c_specification)C Specification

Information about the buffer to get descriptor buffer capture data for is passed in a `VkBufferCaptureDescriptorDataInfoEXT` structure:

```c++
// Provided by VK_EXT_descriptor_buffer
typedef struct VkBufferCaptureDescriptorDataInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkBuffer           buffer;
} VkBufferCaptureDescriptorDataInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `buffer` is the `VkBuffer` handle of the buffer to get opaque capture data for.

## [](#_description)Description

Valid Usage

- [](#VUID-VkBufferCaptureDescriptorDataInfoEXT-buffer-08075)VUID-VkBufferCaptureDescriptorDataInfoEXT-buffer-08075  
  `buffer` **must** have been created with `VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` set in [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html)::`flags`

Valid Usage (Implicit)

- [](#VUID-VkBufferCaptureDescriptorDataInfoEXT-sType-sType)VUID-VkBufferCaptureDescriptorDataInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_CAPTURE_DESCRIPTOR_DATA_INFO_EXT`
- [](#VUID-VkBufferCaptureDescriptorDataInfoEXT-pNext-pNext)VUID-VkBufferCaptureDescriptorDataInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkBufferCaptureDescriptorDataInfoEXT-buffer-parameter)VUID-VkBufferCaptureDescriptorDataInfoEXT-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetBufferOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferOpaqueCaptureDescriptorDataEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferCaptureDescriptorDataInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0