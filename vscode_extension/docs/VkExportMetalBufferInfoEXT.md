# VkExportMetalBufferInfoEXT(3) Manual Page

## Name

VkExportMetalBufferInfoEXT - Structure that identifies a VkDeviceMemory object and corresponding Metal MTLBuffer object



## [](#_c_specification)C Specification

To export the Metal `MTLBuffer` object underlying a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object, include a `VkExportMetalBufferInfoEXT` structure in the `pNext` chain of the `pMetalObjectsInfo` parameter of a [vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkExportMetalObjectsEXT.html) call.

The `VkExportMetalBufferInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_metal_objects
typedef struct VkExportMetalBufferInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceMemory     memory;
    MTLBuffer_id       mtlBuffer;
} VkExportMetalBufferInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `memory` is a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html).
- `mtlBuffer` is the Metal `id<MTLBuffer>` object underlying the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object in `memory`. The implementation will return the `MTLBuffer` in this member, or it will return `NULL` if no `MTLBuffer` could be found underlying the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkExportMetalBufferInfoEXT-sType-sType)VUID-VkExportMetalBufferInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_METAL_BUFFER_INFO_EXT`
- [](#VUID-VkExportMetalBufferInfoEXT-memory-parameter)VUID-VkExportMetalBufferInfoEXT-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle

## [](#_see_also)See Also

[VK\_EXT\_metal\_objects](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_objects.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExportMetalBufferInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0