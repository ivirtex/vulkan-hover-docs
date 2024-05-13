# VkImportMemoryZirconHandleInfoFUCHSIA(3) Manual Page

## Name

VkImportMemoryZirconHandleInfoFUCHSIA - Structure specifying import
parameters for Zircon handle to external memory



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImportMemoryZirconHandleInfoFUCHSIA` structure is defined as:

``` c
// Provided by VK_FUCHSIA_external_memory
typedef struct VkImportMemoryZirconHandleInfoFUCHSIA {
    VkStructureType                       sType;
    const void*                           pNext;
    VkExternalMemoryHandleTypeFlagBits    handleType;
    zx_handle_t                           handle;
} VkImportMemoryZirconHandleInfoFUCHSIA;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `handleType` is a
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value specifying the type of `handle`.

- `handle` is a `zx_handle_t` (Zircon) handle to the external memory.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkImportMemoryZirconHandleInfoFUCHSIA-handleType-04771"
  id="VUID-VkImportMemoryZirconHandleInfoFUCHSIA-handleType-04771"></a>
  VUID-VkImportMemoryZirconHandleInfoFUCHSIA-handleType-04771  
  `handleType` **must** be
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ZIRCON_VMO_BIT_FUCHSIA`

- <a href="#VUID-VkImportMemoryZirconHandleInfoFUCHSIA-handle-04772"
  id="VUID-VkImportMemoryZirconHandleInfoFUCHSIA-handle-04772"></a>
  VUID-VkImportMemoryZirconHandleInfoFUCHSIA-handle-04772  
  `handle` **must** be a valid VMO handle

Valid Usage (Implicit)

- <a href="#VUID-VkImportMemoryZirconHandleInfoFUCHSIA-sType-sType"
  id="VUID-VkImportMemoryZirconHandleInfoFUCHSIA-sType-sType"></a>
  VUID-VkImportMemoryZirconHandleInfoFUCHSIA-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMPORT_MEMORY_ZIRCON_HANDLE_INFO_FUCHSIA`

- <a
  href="#VUID-VkImportMemoryZirconHandleInfoFUCHSIA-handleType-parameter"
  id="VUID-VkImportMemoryZirconHandleInfoFUCHSIA-handleType-parameter"></a>
  VUID-VkImportMemoryZirconHandleInfoFUCHSIA-handleType-parameter  
  If `handleType` is not `0`, `handleType` **must** be a valid
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_external_memory.html),
[VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImportMemoryZirconHandleInfoFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
