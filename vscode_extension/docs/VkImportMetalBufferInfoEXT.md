# VkImportMetalBufferInfoEXT(3) Manual Page

## Name

VkImportMetalBufferInfoEXT - Structure that identifies a Metal MTLBuffer object to use when creating a VkDeviceMemory object.



## [](#_c_specification)C Specification

To import a Metal `MTLBuffer` object to underlie a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object, include a `VkImportMetalBufferInfoEXT` structure in the `pNext` chain of the [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) structure in a [vkAllocateMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAllocateMemory.html) command.

The `VkImportMetalBufferInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_metal_objects
typedef struct VkImportMetalBufferInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    MTLBuffer_id       mtlBuffer;
} VkImportMetalBufferInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `mtlBuffer` is the Metal `id<MTLBuffer>` object that is to underlie the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html).

## [](#_description)Description

The application **must** ensure that the configuration of the `id<MTLBuffer>` object is compatible with the configuration of the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html). Failure to do so results in undefined behavior.

Valid Usage (Implicit)

- [](#VUID-VkImportMetalBufferInfoEXT-sType-sType)VUID-VkImportMetalBufferInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_METAL_BUFFER_INFO_EXT`

## [](#_see_also)See Also

[VK\_EXT\_metal\_objects](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_objects.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImportMetalBufferInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0