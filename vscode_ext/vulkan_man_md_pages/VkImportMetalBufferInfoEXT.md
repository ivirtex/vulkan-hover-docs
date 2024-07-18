# VkImportMetalBufferInfoEXT(3) Manual Page

## Name

VkImportMetalBufferInfoEXT - Structure that identifies a Metal MTLBuffer
object to use when creating a VkDeviceMemory object.



## <a href="#_c_specification" class="anchor"></a>C Specification

To import a Metal `MTLBuffer` object to underlie a
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object, include a
`VkImportMetalBufferInfoEXT` structure in the `pNext` chain of the
[VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) structure in a
[vkAllocateMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAllocateMemory.html) command.

The `VkImportMetalBufferInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_metal_objects
typedef struct VkImportMetalBufferInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    MTLBuffer_id       mtlBuffer;
} VkImportMetalBufferInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `mtlBuffer` is the Metal `id<MTLBuffer>` object that is to underlie
  the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html).

## <a href="#_description" class="anchor"></a>Description

The application **must** ensure that the configuration of the
`id<MTLBuffer>` object is compatible with the configuration of the
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html). Failure to do so results in
undefined behavior.

Valid Usage (Implicit)

- <a href="#VUID-VkImportMetalBufferInfoEXT-sType-sType"
  id="VUID-VkImportMetalBufferInfoEXT-sType-sType"></a>
  VUID-VkImportMetalBufferInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_METAL_BUFFER_INFO_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_objects](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_objects.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImportMetalBufferInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
