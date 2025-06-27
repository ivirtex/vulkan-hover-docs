# VkImportSemaphoreZirconHandleInfoFUCHSIA(3) Manual Page

## Name

VkImportSemaphoreZirconHandleInfoFUCHSIA - Structure specifying Zircon
event handle to import to a semaphore



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImportSemaphoreZirconHandleInfoFUCHSIA` structure is defined as:

``` c
// Provided by VK_FUCHSIA_external_semaphore
typedef struct VkImportSemaphoreZirconHandleInfoFUCHSIA {
    VkStructureType                          sType;
    const void*                              pNext;
    VkSemaphore                              semaphore;
    VkSemaphoreImportFlags                   flags;
    VkExternalSemaphoreHandleTypeFlagBits    handleType;
    zx_handle_t                              zirconHandle;
} VkImportSemaphoreZirconHandleInfoFUCHSIA;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `semaphore` is the semaphore into which the payload will be imported.

- `flags` is a bitmask of
  [VkSemaphoreImportFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlagBits.html) specifying
  additional parameters for the semaphore payload import operation.

- `handleType` is a
  [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html)
  value specifying the type of `zirconHandle`.

- `zirconHandle` is the external handle to import.

## <a href="#_description" class="anchor"></a>Description

The handle types supported by `handleType` are:

| Handle Type | Transference | Permanence Supported |
|----|----|----|
| `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_ZIRCON_EVENT_BIT_FUCHSIA` | Reference | Temporary,Permanent |

Table 1. Handle Types Supported by
`VkImportSemaphoreZirconHandleInfoFUCHSIA`
{#synchronization-semaphore-handletypes-fuchsia}

Valid Usage

- <a
  href="#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-handleType-04765"
  id="VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-handleType-04765"></a>
  VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-handleType-04765  
  `handleType` **must** be a value included in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphore-handletypes-fuchsia"
  target="_blank" rel="noopener">Handle Types Supported by
  <code>VkImportSemaphoreZirconHandleInfoFUCHSIA</code></a> table

- <a
  href="#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-zirconHandle-04766"
  id="VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-zirconHandle-04766"></a>
  VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-zirconHandle-04766  
  `zirconHandle` **must** obey any requirements listed for `handleType`
  in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#external-semaphore-handle-types-compatibility"
  target="_blank" rel="noopener">external semaphore handle types
  compatibility</a>

- <a
  href="#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-zirconHandle-04767"
  id="VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-zirconHandle-04767"></a>
  VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-zirconHandle-04767  
  `zirconHandle` **must** have `ZX_RIGHTS_BASIC` and `ZX_RIGHTS_SIGNAL`
  rights

- <a
  href="#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-semaphoreType-04768"
  id="VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-semaphoreType-04768"></a>
  VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-semaphoreType-04768  
  The
  [VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeCreateInfo.html)::`semaphoreType`
  field **must** not be `VK_SEMAPHORE_TYPE_TIMELINE`

Valid Usage (Implicit)

- <a href="#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-sType-sType"
  id="VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-sType-sType"></a>
  VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_ZIRCON_HANDLE_INFO_FUCHSIA`

- <a href="#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-pNext-pNext"
  id="VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-pNext-pNext"></a>
  VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-semaphore-parameter"
  id="VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-semaphore-parameter"></a>
  VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-semaphore-parameter  
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handle

- <a href="#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-flags-parameter"
  id="VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-flags-parameter"></a>
  VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-flags-parameter  
  `flags` **must** be a valid combination of
  [VkSemaphoreImportFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlagBits.html) values

- <a
  href="#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-handleType-parameter"
  id="VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-handleType-parameter"></a>
  VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html)
  value

Host Synchronization

- Host access to `semaphore` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_external_semaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_external_semaphore.html),
[VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html),
[VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html),
[VkSemaphoreImportFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkImportSemaphoreZirconHandleFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkImportSemaphoreZirconHandleFUCHSIA.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImportSemaphoreZirconHandleInfoFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
