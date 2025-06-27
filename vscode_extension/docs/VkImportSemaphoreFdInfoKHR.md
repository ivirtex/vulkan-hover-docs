# VkImportSemaphoreFdInfoKHR(3) Manual Page

## Name

VkImportSemaphoreFdInfoKHR - Structure specifying POSIX file descriptor
to import to a semaphore



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImportSemaphoreFdInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_external_semaphore_fd
typedef struct VkImportSemaphoreFdInfoKHR {
    VkStructureType                          sType;
    const void*                              pNext;
    VkSemaphore                              semaphore;
    VkSemaphoreImportFlags                   flags;
    VkExternalSemaphoreHandleTypeFlagBits    handleType;
    int                                      fd;
} VkImportSemaphoreFdInfoKHR;
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
  value specifying the type of `fd`.

- `fd` is the external handle to import.

## <a href="#_description" class="anchor"></a>Description

The handle types supported by `handleType` are:

| Handle Type | Transference | Permanence Supported |
|----|----|----|
| `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT` | Reference | Temporary,Permanent |
| `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT` | Copy | Temporary |

Table 1. Handle Types Supported by `VkImportSemaphoreFdInfoKHR`
{#synchronization-semaphore-handletypes-fd}

Valid Usage

- <a href="#VUID-VkImportSemaphoreFdInfoKHR-handleType-01143"
  id="VUID-VkImportSemaphoreFdInfoKHR-handleType-01143"></a>
  VUID-VkImportSemaphoreFdInfoKHR-handleType-01143  
  `handleType` **must** be a value included in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphore-handletypes-fd"
  target="_blank" rel="noopener">Handle Types Supported by
  <code>VkImportSemaphoreFdInfoKHR</code></a> table

- <a href="#VUID-VkImportSemaphoreFdInfoKHR-fd-01544"
  id="VUID-VkImportSemaphoreFdInfoKHR-fd-01544"></a>
  VUID-VkImportSemaphoreFdInfoKHR-fd-01544  
  `fd` **must** obey any requirements listed for `handleType` in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#external-semaphore-handle-types-compatibility"
  target="_blank" rel="noopener">external semaphore handle types
  compatibility</a>

- <a href="#VUID-VkImportSemaphoreFdInfoKHR-handleType-03263"
  id="VUID-VkImportSemaphoreFdInfoKHR-handleType-03263"></a>
  VUID-VkImportSemaphoreFdInfoKHR-handleType-03263  
  If `handleType` is `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT`,
  the [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html)::`flags` field
  **must** match that of the semaphore from which `fd` was exported

- <a href="#VUID-VkImportSemaphoreFdInfoKHR-handleType-07307"
  id="VUID-VkImportSemaphoreFdInfoKHR-handleType-07307"></a>
  VUID-VkImportSemaphoreFdInfoKHR-handleType-07307  
  If `handleType` refers to a handle type with copy payload transference
  semantics, `flags` **must** contain
  `VK_SEMAPHORE_IMPORT_TEMPORARY_BIT`

- <a href="#VUID-VkImportSemaphoreFdInfoKHR-handleType-03264"
  id="VUID-VkImportSemaphoreFdInfoKHR-handleType-03264"></a>
  VUID-VkImportSemaphoreFdInfoKHR-handleType-03264  
  If `handleType` is `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT`,
  the
  [VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeCreateInfo.html)::`semaphoreType`
  field **must** match that of the semaphore from which `fd` was
  exported

- <a href="#VUID-VkImportSemaphoreFdInfoKHR-flags-03323"
  id="VUID-VkImportSemaphoreFdInfoKHR-flags-03323"></a>
  VUID-VkImportSemaphoreFdInfoKHR-flags-03323  
  If `flags` contains `VK_SEMAPHORE_IMPORT_TEMPORARY_BIT`, the
  [VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeCreateInfo.html)::`semaphoreType`
  field of the semaphore from which `fd` was exported **must** not be
  `VK_SEMAPHORE_TYPE_TIMELINE`

If `handleType` is `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT`, the
special value `-1` for `fd` is treated like a valid sync file descriptor
referring to an object that has already signaled. The import operation
will succeed and the `VkSemaphore` will have a temporarily imported
payload as if a valid file descriptor had been provided.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>This special behavior for importing an invalid sync file descriptor
allows easier interoperability with other system APIs which use the
convention that an invalid sync file descriptor represents work that has
already completed and does not need to be waited for. It is consistent
with the option for implementations to return a <code>-1</code> file
descriptor when exporting a
<code>VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT</code> from a
<code>VkSemaphore</code> which is signaled.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkImportSemaphoreFdInfoKHR-sType-sType"
  id="VUID-VkImportSemaphoreFdInfoKHR-sType-sType"></a>
  VUID-VkImportSemaphoreFdInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_FD_INFO_KHR`

- <a href="#VUID-VkImportSemaphoreFdInfoKHR-pNext-pNext"
  id="VUID-VkImportSemaphoreFdInfoKHR-pNext-pNext"></a>
  VUID-VkImportSemaphoreFdInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkImportSemaphoreFdInfoKHR-semaphore-parameter"
  id="VUID-VkImportSemaphoreFdInfoKHR-semaphore-parameter"></a>
  VUID-VkImportSemaphoreFdInfoKHR-semaphore-parameter  
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handle

- <a href="#VUID-VkImportSemaphoreFdInfoKHR-flags-parameter"
  id="VUID-VkImportSemaphoreFdInfoKHR-flags-parameter"></a>
  VUID-VkImportSemaphoreFdInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of
  [VkSemaphoreImportFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlagBits.html) values

- <a href="#VUID-VkImportSemaphoreFdInfoKHR-handleType-parameter"
  id="VUID-VkImportSemaphoreFdInfoKHR-handleType-parameter"></a>
  VUID-VkImportSemaphoreFdInfoKHR-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html)
  value

Host Synchronization

- Host access to `semaphore` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_semaphore_fd](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_semaphore_fd.html),
[VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html),
[VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html),
[VkSemaphoreImportFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkImportSemaphoreFdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkImportSemaphoreFdKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImportSemaphoreFdInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
