# VkImportFenceFdInfoKHR(3) Manual Page

## Name

VkImportFenceFdInfoKHR - (None)



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImportFenceFdInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_external_fence_fd
typedef struct VkImportFenceFdInfoKHR {
    VkStructureType                      sType;
    const void*                          pNext;
    VkFence                              fence;
    VkFenceImportFlags                   flags;
    VkExternalFenceHandleTypeFlagBits    handleType;
    int                                  fd;
} VkImportFenceFdInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `fence` is the fence into which the payload will be imported.

- `flags` is a bitmask of
  [VkFenceImportFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceImportFlagBits.html) specifying
  additional parameters for the fence payload import operation.

- `handleType` is a
  [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html)
  value specifying the type of `fd`.

- `fd` is the external handle to import.

## <a href="#_description" class="anchor"></a>Description

The handle types supported by `handleType` are:

| Handle Type | Transference | Permanence Supported |
|----|----|----|
| `VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT` | Reference | Temporary,Permanent |
| `VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT` | Copy | Temporary |

Table 1. Handle Types Supported by `VkImportFenceFdInfoKHR`
{#synchronization-fence-handletypes-fd}

Valid Usage

- <a href="#VUID-VkImportFenceFdInfoKHR-handleType-01464"
  id="VUID-VkImportFenceFdInfoKHR-handleType-01464"></a>
  VUID-VkImportFenceFdInfoKHR-handleType-01464  
  `handleType` **must** be a value included in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-fence-handletypes-fd"
  target="_blank" rel="noopener">Handle Types Supported by
  <code>VkImportFenceFdInfoKHR</code></a> table

- <a href="#VUID-VkImportFenceFdInfoKHR-fd-01541"
  id="VUID-VkImportFenceFdInfoKHR-fd-01541"></a>
  VUID-VkImportFenceFdInfoKHR-fd-01541  
  `fd` **must** obey any requirements listed for `handleType` in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#external-fence-handle-types-compatibility"
  target="_blank" rel="noopener">external fence handle types
  compatibility</a>

- <a href="#VUID-VkImportFenceFdInfoKHR-handleType-07306"
  id="VUID-VkImportFenceFdInfoKHR-handleType-07306"></a>
  VUID-VkImportFenceFdInfoKHR-handleType-07306  
  If `handleType` refers to a handle type with copy payload transference
  semantics, `flags` **must** contain `VK_FENCE_IMPORT_TEMPORARY_BIT`

If `handleType` is `VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT`, the
special value `-1` for `fd` is treated like a valid sync file descriptor
referring to an object that has already signaled. The import operation
will succeed and the `VkFence` will have a temporarily imported payload
as if a valid file descriptor had been provided.

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
<code>VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT</code> from a
<code>VkFence</code> which is signaled.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkImportFenceFdInfoKHR-sType-sType"
  id="VUID-VkImportFenceFdInfoKHR-sType-sType"></a>
  VUID-VkImportFenceFdInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_FENCE_FD_INFO_KHR`

- <a href="#VUID-VkImportFenceFdInfoKHR-pNext-pNext"
  id="VUID-VkImportFenceFdInfoKHR-pNext-pNext"></a>
  VUID-VkImportFenceFdInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkImportFenceFdInfoKHR-fence-parameter"
  id="VUID-VkImportFenceFdInfoKHR-fence-parameter"></a>
  VUID-VkImportFenceFdInfoKHR-fence-parameter  
  `fence` **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html) handle

- <a href="#VUID-VkImportFenceFdInfoKHR-flags-parameter"
  id="VUID-VkImportFenceFdInfoKHR-flags-parameter"></a>
  VUID-VkImportFenceFdInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of
  [VkFenceImportFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceImportFlagBits.html) values

- <a href="#VUID-VkImportFenceFdInfoKHR-handleType-parameter"
  id="VUID-VkImportFenceFdInfoKHR-handleType-parameter"></a>
  VUID-VkImportFenceFdInfoKHR-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html)
  value

Host Synchronization

- Host access to `fence` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_fence_fd](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_fence_fd.html),
[VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html),
[VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html), [VkFenceImportFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceImportFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkImportFenceFdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkImportFenceFdKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImportFenceFdInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
