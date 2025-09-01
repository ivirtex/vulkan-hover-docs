# VkImportScreenBufferInfoQNX(3) Manual Page

## Name

VkImportScreenBufferInfoQNX - Import memory from a QNX Screen buffer



## [](#_c_specification)C Specification

To import memory created outside of the current Vulkan instance from a QNX Screen buffer, add a `VkImportScreenBufferInfoQNX` structure to the `pNext` chain of the [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) structure. The `VkImportScreenBufferInfoQNX` structure is defined as:

```c++
// Provided by VK_QNX_external_memory_screen_buffer
typedef struct VkImportScreenBufferInfoQNX {
    VkStructureType           sType;
    const void*               pNext;
    struct _screen_buffer*    buffer;
} VkImportScreenBufferInfoQNX;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `buffer` is a pointer to a `struct` `_screen_buffer`, the QNX Screen buffer to import

## [](#_description)Description

The implementation **may** not acquire a reference to the imported Screen buffer. Therefore, the application **must** ensure that the object referred to by `buffer` stays valid as long as the device memory to which it is imported is being used.

Valid Usage

- [](#VUID-VkImportScreenBufferInfoQNX-buffer-08966)VUID-VkImportScreenBufferInfoQNX-buffer-08966  
  If `buffer` is not `NULL`, QNX Screen Buffers **must** be supported for import, as reported by [VkExternalImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatProperties.html) or [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalBufferProperties.html)
- [](#VUID-VkImportScreenBufferInfoQNX-buffer-08967)VUID-VkImportScreenBufferInfoQNX-buffer-08967  
  `buffer` is not `NULL`, it **must** be a pointer to [valid QNX Screen buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-external-screen-buffer-validity)

Valid Usage (Implicit)

- [](#VUID-VkImportScreenBufferInfoQNX-sType-sType)VUID-VkImportScreenBufferInfoQNX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_SCREEN_BUFFER_INFO_QNX`

## [](#_see_also)See Also

[VK\_QNX\_external\_memory\_screen\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QNX_external_memory_screen_buffer.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImportScreenBufferInfoQNX).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0