# VkMemoryGetZirconHandleInfoFUCHSIA(3) Manual Page

## Name

VkMemoryGetZirconHandleInfoFUCHSIA - Structure specifying export
parameters for Zircon handle to device memory



## <a href="#_c_specification" class="anchor"></a>C Specification

`VkMemoryGetZirconHandleInfoFUCHSIA` is defined as:

``` c
// Provided by VK_FUCHSIA_external_memory
typedef struct VkMemoryGetZirconHandleInfoFUCHSIA {
    VkStructureType                       sType;
    const void*                           pNext;
    VkDeviceMemory                        memory;
    VkExternalMemoryHandleTypeFlagBits    handleType;
} VkMemoryGetZirconHandleInfoFUCHSIA;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `memory` the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) being exported.

- `handleType` is a
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value specifying the type of the handle pointed to by
  [vkGetMemoryZirconHandleFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryZirconHandleFUCHSIA.html)::`pZirconHandle`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkMemoryGetZirconHandleInfoFUCHSIA-handleType-04775"
  id="VUID-VkMemoryGetZirconHandleInfoFUCHSIA-handleType-04775"></a>
  VUID-VkMemoryGetZirconHandleInfoFUCHSIA-handleType-04775  
  `handleType` **must** be
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ZIRCON_VMO_BIT_FUCHSIA`

- <a href="#VUID-VkMemoryGetZirconHandleInfoFUCHSIA-handleType-04776"
  id="VUID-VkMemoryGetZirconHandleInfoFUCHSIA-handleType-04776"></a>
  VUID-VkMemoryGetZirconHandleInfoFUCHSIA-handleType-04776  
  `handleType` **must** have been included in the `handleTypes` field of
  the `VkExportMemoryAllocateInfo` structure when the external memory
  was allocated

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryGetZirconHandleInfoFUCHSIA-sType-sType"
  id="VUID-VkMemoryGetZirconHandleInfoFUCHSIA-sType-sType"></a>
  VUID-VkMemoryGetZirconHandleInfoFUCHSIA-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_MEMORY_GET_ZIRCON_HANDLE_INFO_FUCHSIA`

- <a href="#VUID-VkMemoryGetZirconHandleInfoFUCHSIA-pNext-pNext"
  id="VUID-VkMemoryGetZirconHandleInfoFUCHSIA-pNext-pNext"></a>
  VUID-VkMemoryGetZirconHandleInfoFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkMemoryGetZirconHandleInfoFUCHSIA-memory-parameter"
  id="VUID-VkMemoryGetZirconHandleInfoFUCHSIA-memory-parameter"></a>
  VUID-VkMemoryGetZirconHandleInfoFUCHSIA-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

- <a href="#VUID-VkMemoryGetZirconHandleInfoFUCHSIA-handleType-parameter"
  id="VUID-VkMemoryGetZirconHandleInfoFUCHSIA-handleType-parameter"></a>
  VUID-VkMemoryGetZirconHandleInfoFUCHSIA-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_external_memory.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetMemoryZirconHandleFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryZirconHandleFUCHSIA.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryGetZirconHandleInfoFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
