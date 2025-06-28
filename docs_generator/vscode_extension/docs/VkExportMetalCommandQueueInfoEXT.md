# VkExportMetalCommandQueueInfoEXT(3) Manual Page

## Name

VkExportMetalCommandQueueInfoEXT - Structure that identifies a VkQueue object and corresponding Metal MTLCommandQueue object



## [](#_c_specification)C Specification

To export the Metal `MTLCommandQueue` object underlying a [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) object, include a `VkExportMetalCommandQueueInfoEXT` structure in the `pNext` chain of the `pMetalObjectsInfo` parameter of a [vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkExportMetalObjectsEXT.html) call.

The `VkExportMetalCommandQueueInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_metal_objects
typedef struct VkExportMetalCommandQueueInfoEXT {
    VkStructureType       sType;
    const void*           pNext;
    VkQueue               queue;
    MTLCommandQueue_id    mtlCommandQueue;
} VkExportMetalCommandQueueInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `queue` is a [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html).
- `mtlCommandQueue` is the Metal `id<MTLCommandQueue>` object underlying the [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) object in `queue`. The implementation will return the `MTLCommandQueue` in this member, or it will return `NULL` if no `MTLCommandQueue` could be found underlying the [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) object.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkExportMetalCommandQueueInfoEXT-sType-sType)VUID-VkExportMetalCommandQueueInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_METAL_COMMAND_QUEUE_INFO_EXT`
- [](#VUID-VkExportMetalCommandQueueInfoEXT-queue-parameter)VUID-VkExportMetalCommandQueueInfoEXT-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) handle

## [](#_see_also)See Also

[VK\_EXT\_metal\_objects](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_objects.html), [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExportMetalCommandQueueInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0