# VkImportMemoryZirconHandleInfoFUCHSIA(3) Manual Page

## Name

VkImportMemoryZirconHandleInfoFUCHSIA - Structure specifying import parameters for Zircon handle to external memory



## [](#_c_specification)C Specification

The `VkImportMemoryZirconHandleInfoFUCHSIA` structure is defined as:

```c++
// Provided by VK_FUCHSIA_external_memory
typedef struct VkImportMemoryZirconHandleInfoFUCHSIA {
    VkStructureType                       sType;
    const void*                           pNext;
    VkExternalMemoryHandleTypeFlagBits    handleType;
    zx_handle_t                           handle;
} VkImportMemoryZirconHandleInfoFUCHSIA;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `handleType` is a [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value specifying the type of `handle`.
- `handle` is a `zx_handle_t` (Zircon) handle to the external memory.

## [](#_description)Description

Valid Usage

- [](#VUID-VkImportMemoryZirconHandleInfoFUCHSIA-handleType-04771)VUID-VkImportMemoryZirconHandleInfoFUCHSIA-handleType-04771  
  `handleType` **must** be `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ZIRCON_VMO_BIT_FUCHSIA`
- [](#VUID-VkImportMemoryZirconHandleInfoFUCHSIA-handle-04772)VUID-VkImportMemoryZirconHandleInfoFUCHSIA-handle-04772  
  `handle` **must** be a valid VMO handle

Valid Usage (Implicit)

- [](#VUID-VkImportMemoryZirconHandleInfoFUCHSIA-sType-sType)VUID-VkImportMemoryZirconHandleInfoFUCHSIA-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_MEMORY_ZIRCON_HANDLE_INFO_FUCHSIA`
- [](#VUID-VkImportMemoryZirconHandleInfoFUCHSIA-handleType-parameter)VUID-VkImportMemoryZirconHandleInfoFUCHSIA-handleType-parameter  
  If `handleType` is not `0`, `handleType` **must** be a valid [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value

## [](#_see_also)See Also

[VK\_FUCHSIA\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_external_memory.html), [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImportMemoryZirconHandleInfoFUCHSIA)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0