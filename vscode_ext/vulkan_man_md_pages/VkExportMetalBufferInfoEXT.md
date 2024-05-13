# VkExportMetalBufferInfoEXT(3) Manual Page

## Name

VkExportMetalBufferInfoEXT - Structure that identifies a VkDeviceMemory
object and corresponding Metal MTLBuffer object



## <a href="#_c_specification" class="anchor"></a>C Specification

To export the Metal `MTLBuffer` object underlying a
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object, include a
`VkExportMetalBufferInfoEXT` structure in the `pNext` chain of the
`pMetalObjectsInfo` parameter of a
[vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkExportMetalObjectsEXT.html) call.

The `VkExportMetalBufferInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_metal_objects
typedef struct VkExportMetalBufferInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceMemory     memory;
    MTLBuffer_id       mtlBuffer;
} VkExportMetalBufferInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `memory` is a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html).

- `mtlBuffer` is the Metal `id<MTLBuffer>` object underlying the
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object in `memory`. The
  implementation will return the `MTLBuffer` in this member, or it will
  return `NULL` if no `MTLBuffer` could be found underlying the
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkExportMetalBufferInfoEXT-sType-sType"
  id="VUID-VkExportMetalBufferInfoEXT-sType-sType"></a>
  VUID-VkExportMetalBufferInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_METAL_BUFFER_INFO_EXT`

- <a href="#VUID-VkExportMetalBufferInfoEXT-memory-parameter"
  id="VUID-VkExportMetalBufferInfoEXT-memory-parameter"></a>
  VUID-VkExportMetalBufferInfoEXT-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_objects](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_objects.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExportMetalBufferInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
