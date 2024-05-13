# VkImportSemaphoreWin32HandleInfoKHR(3) Manual Page

## Name

VkImportSemaphoreWin32HandleInfoKHR - Structure specifying Windows
handle to import to a semaphore



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImportSemaphoreWin32HandleInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_external_semaphore_win32
typedef struct VkImportSemaphoreWin32HandleInfoKHR {
    VkStructureType                          sType;
    const void*                              pNext;
    VkSemaphore                              semaphore;
    VkSemaphoreImportFlags                   flags;
    VkExternalSemaphoreHandleTypeFlagBits    handleType;
    HANDLE                                   handle;
    LPCWSTR                                  name;
} VkImportSemaphoreWin32HandleInfoKHR;
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
  value specifying the type of `handle`.

- `handle` is `NULL` or the external handle to import.

- `name` is `NULL` or a null-terminated UTF-16 string naming the
  underlying synchronization primitive to import.

## <a href="#_description" class="anchor"></a>Description

The handle types supported by `handleType` are:

| Handle Type                                              | Transference | Permanence Supported |
|----------------------------------------------------------|--------------|----------------------|
| `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT`     | Reference    | Temporary,Permanent  |
| `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT` | Reference    | Temporary,Permanent  |
| `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT`      | Reference    | Temporary,Permanent  |

Table 1. Handle Types Supported by `VkImportSemaphoreWin32HandleInfoKHR`
{#synchronization-semaphore-handletypes-win32}

Valid Usage

- <a href="#VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01140"
  id="VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01140"></a>
  VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01140  
  `handleType` **must** be a value included in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphore-handletypes-win32"
  target="_blank" rel="noopener">Handle Types Supported by
  <code>VkImportSemaphoreWin32HandleInfoKHR</code></a> table

- <a href="#VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01466"
  id="VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01466"></a>
  VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01466  
  If `handleType` is not
  `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT` or
  `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT`, `name` **must**
  be `NULL`

- <a href="#VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01467"
  id="VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01467"></a>
  VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01467  
  If `handle` is `NULL`, `name` **must** name a valid synchronization
  primitive of the type specified by `handleType`

- <a href="#VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01468"
  id="VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01468"></a>
  VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01468  
  If `name` is `NULL`, `handle` **must** be a valid handle of the type
  specified by `handleType`

- <a href="#VUID-VkImportSemaphoreWin32HandleInfoKHR-handle-01469"
  id="VUID-VkImportSemaphoreWin32HandleInfoKHR-handle-01469"></a>
  VUID-VkImportSemaphoreWin32HandleInfoKHR-handle-01469  
  If `handle` is not `NULL`, `name` **must** be `NULL`

- <a href="#VUID-VkImportSemaphoreWin32HandleInfoKHR-handle-01542"
  id="VUID-VkImportSemaphoreWin32HandleInfoKHR-handle-01542"></a>
  VUID-VkImportSemaphoreWin32HandleInfoKHR-handle-01542  
  If `handle` is not `NULL`, it **must** obey any requirements listed
  for `handleType` in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#external-semaphore-handle-types-compatibility"
  target="_blank" rel="noopener">external semaphore handle types
  compatibility</a>

- <a href="#VUID-VkImportSemaphoreWin32HandleInfoKHR-name-01543"
  id="VUID-VkImportSemaphoreWin32HandleInfoKHR-name-01543"></a>
  VUID-VkImportSemaphoreWin32HandleInfoKHR-name-01543  
  If `name` is not `NULL`, it **must** obey any requirements listed for
  `handleType` in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#external-semaphore-handle-types-compatibility"
  target="_blank" rel="noopener">external semaphore handle types
  compatibility</a>

- <a href="#VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-03261"
  id="VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-03261"></a>
  VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-03261  
  If `handleType` is
  `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT` or
  `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT`, the
  [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html)::`flags` field
  **must** match that of the semaphore from which `handle` or `name` was
  exported

- <a href="#VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-03262"
  id="VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-03262"></a>
  VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-03262  
  If `handleType` is
  `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT` or
  `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT`, the
  [VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeCreateInfo.html)::`semaphoreType`
  field **must** match that of the semaphore from which `handle` or
  `name` was exported

- <a href="#VUID-VkImportSemaphoreWin32HandleInfoKHR-flags-03322"
  id="VUID-VkImportSemaphoreWin32HandleInfoKHR-flags-03322"></a>
  VUID-VkImportSemaphoreWin32HandleInfoKHR-flags-03322  
  If `flags` contains `VK_SEMAPHORE_IMPORT_TEMPORARY_BIT`, the
  [VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeCreateInfo.html)::`semaphoreType`
  field of the semaphore from which `handle` or `name` was exported
  **must** not be `VK_SEMAPHORE_TYPE_TIMELINE`

Valid Usage (Implicit)

- <a href="#VUID-VkImportSemaphoreWin32HandleInfoKHR-sType-sType"
  id="VUID-VkImportSemaphoreWin32HandleInfoKHR-sType-sType"></a>
  VUID-VkImportSemaphoreWin32HandleInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR`

- <a href="#VUID-VkImportSemaphoreWin32HandleInfoKHR-pNext-pNext"
  id="VUID-VkImportSemaphoreWin32HandleInfoKHR-pNext-pNext"></a>
  VUID-VkImportSemaphoreWin32HandleInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkImportSemaphoreWin32HandleInfoKHR-semaphore-parameter"
  id="VUID-VkImportSemaphoreWin32HandleInfoKHR-semaphore-parameter"></a>
  VUID-VkImportSemaphoreWin32HandleInfoKHR-semaphore-parameter  
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handle

- <a href="#VUID-VkImportSemaphoreWin32HandleInfoKHR-flags-parameter"
  id="VUID-VkImportSemaphoreWin32HandleInfoKHR-flags-parameter"></a>
  VUID-VkImportSemaphoreWin32HandleInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of
  [VkSemaphoreImportFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlagBits.html) values

Host Synchronization

- Host access to `semaphore` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_semaphore_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_semaphore_win32.html),
[VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html),
[VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html),
[VkSemaphoreImportFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkImportSemaphoreWin32HandleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkImportSemaphoreWin32HandleKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImportSemaphoreWin32HandleInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
