# VkExportMetalCommandQueueInfoEXT(3) Manual Page

## Name

VkExportMetalCommandQueueInfoEXT - Structure that identifies a VkQueue
object and corresponding Metal MTLCommandQueue object



## <a href="#_c_specification" class="anchor"></a>C Specification

To export the Metal `MTLCommandQueue` object underlying a
[VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) object, include a
`VkExportMetalCommandQueueInfoEXT` structure in the `pNext` chain of the
`pMetalObjectsInfo` parameter of a
[vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkExportMetalObjectsEXT.html) call.

The `VkExportMetalCommandQueueInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_metal_objects
typedef struct VkExportMetalCommandQueueInfoEXT {
    VkStructureType       sType;
    const void*           pNext;
    VkQueue               queue;
    MTLCommandQueue_id    mtlCommandQueue;
} VkExportMetalCommandQueueInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `queue` is a [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html).

- `mtlCommandQueue` is the Metal `id<MTLCommandQueue>` object underlying
  the [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) object in `queue`. The implementation will
  return the `MTLCommandQueue` in this member, or it will return `NULL`
  if no `MTLCommandQueue` could be found underlying the
  [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) object.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkExportMetalCommandQueueInfoEXT-sType-sType"
  id="VUID-VkExportMetalCommandQueueInfoEXT-sType-sType"></a>
  VUID-VkExportMetalCommandQueueInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_EXPORT_METAL_COMMAND_QUEUE_INFO_EXT`

- <a href="#VUID-VkExportMetalCommandQueueInfoEXT-queue-parameter"
  id="VUID-VkExportMetalCommandQueueInfoEXT-queue-parameter"></a>
  VUID-VkExportMetalCommandQueueInfoEXT-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_objects](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_objects.html),
[VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExportMetalCommandQueueInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
