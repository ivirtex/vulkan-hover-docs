# VkMemoryZirconHandlePropertiesFUCHSIA(3) Manual Page

## Name

VkMemoryZirconHandlePropertiesFUCHSIA - Structure specifying Zircon
handle compatible external memory



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMemoryZirconHandlePropertiesFUCHSIA` structure is defined as:

``` c
// Provided by VK_FUCHSIA_external_memory
typedef struct VkMemoryZirconHandlePropertiesFUCHSIA {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           memoryTypeBits;
} VkMemoryZirconHandlePropertiesFUCHSIA;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `memoryTypeBits` a bitmask containing one bit set for every memory
  type which the specified handle can be imported as.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryZirconHandlePropertiesFUCHSIA-sType-sType"
  id="VUID-VkMemoryZirconHandlePropertiesFUCHSIA-sType-sType"></a>
  VUID-VkMemoryZirconHandlePropertiesFUCHSIA-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_MEMORY_ZIRCON_HANDLE_PROPERTIES_FUCHSIA`

- <a href="#VUID-VkMemoryZirconHandlePropertiesFUCHSIA-pNext-pNext"
  id="VUID-VkMemoryZirconHandlePropertiesFUCHSIA-pNext-pNext"></a>
  VUID-VkMemoryZirconHandlePropertiesFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_external_memory.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetMemoryZirconHandlePropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryZirconHandlePropertiesFUCHSIA.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryZirconHandlePropertiesFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
