# VkImportFenceWin32HandleInfoKHR(3) Manual Page

## Name

VkImportFenceWin32HandleInfoKHR - (None)



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImportFenceWin32HandleInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_external_fence_win32
typedef struct VkImportFenceWin32HandleInfoKHR {
    VkStructureType                      sType;
    const void*                          pNext;
    VkFence                              fence;
    VkFenceImportFlags                   flags;
    VkExternalFenceHandleTypeFlagBits    handleType;
    HANDLE                               handle;
    LPCWSTR                              name;
} VkImportFenceWin32HandleInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `fence` is the fence into which the state will be imported.

- `flags` is a bitmask of
  [VkFenceImportFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceImportFlagBits.html) specifying
  additional parameters for the fence payload import operation.

- `handleType` is a
  [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html)
  value specifying the type of `handle`.

- `handle` is `NULL` or the external handle to import.

- `name` is `NULL` or a null-terminated UTF-16 string naming the
  underlying synchronization primitive to import.

## <a href="#_description" class="anchor"></a>Description

The handle types supported by `handleType` are:

| Handle Type | Transference | Permanence Supported |
|----|----|----|
| `VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT` | Reference | Temporary,Permanent |
| `VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT` | Reference | Temporary,Permanent |

Table 1. Handle Types Supported by `VkImportFenceWin32HandleInfoKHR`
{#synchronization-fence-handletypes-win32}

Valid Usage

- <a href="#VUID-VkImportFenceWin32HandleInfoKHR-handleType-01457"
  id="VUID-VkImportFenceWin32HandleInfoKHR-handleType-01457"></a>
  VUID-VkImportFenceWin32HandleInfoKHR-handleType-01457  
  `handleType` **must** be a value included in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-fence-handletypes-win32"
  target="_blank" rel="noopener">Handle Types Supported by
  <code>VkImportFenceWin32HandleInfoKHR</code></a> table

- <a href="#VUID-VkImportFenceWin32HandleInfoKHR-handleType-01459"
  id="VUID-VkImportFenceWin32HandleInfoKHR-handleType-01459"></a>
  VUID-VkImportFenceWin32HandleInfoKHR-handleType-01459  
  If `handleType` is not
  `VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT`, `name` **must** be
  `NULL`

- <a href="#VUID-VkImportFenceWin32HandleInfoKHR-handleType-01460"
  id="VUID-VkImportFenceWin32HandleInfoKHR-handleType-01460"></a>
  VUID-VkImportFenceWin32HandleInfoKHR-handleType-01460  
  If `handle` is `NULL`, `name` **must** name a valid synchronization
  primitive of the type specified by `handleType`

- <a href="#VUID-VkImportFenceWin32HandleInfoKHR-handleType-01461"
  id="VUID-VkImportFenceWin32HandleInfoKHR-handleType-01461"></a>
  VUID-VkImportFenceWin32HandleInfoKHR-handleType-01461  
  If `name` is `NULL`, `handle` **must** be a valid handle of the type
  specified by `handleType`

- <a href="#VUID-VkImportFenceWin32HandleInfoKHR-handle-01462"
  id="VUID-VkImportFenceWin32HandleInfoKHR-handle-01462"></a>
  VUID-VkImportFenceWin32HandleInfoKHR-handle-01462  
  If `handle` is not `NULL`, `name` **must** be `NULL`

- <a href="#VUID-VkImportFenceWin32HandleInfoKHR-handle-01539"
  id="VUID-VkImportFenceWin32HandleInfoKHR-handle-01539"></a>
  VUID-VkImportFenceWin32HandleInfoKHR-handle-01539  
  If `handle` is not `NULL`, it **must** obey any requirements listed
  for `handleType` in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#external-fence-handle-types-compatibility"
  target="_blank" rel="noopener">external fence handle types
  compatibility</a>

- <a href="#VUID-VkImportFenceWin32HandleInfoKHR-name-01540"
  id="VUID-VkImportFenceWin32HandleInfoKHR-name-01540"></a>
  VUID-VkImportFenceWin32HandleInfoKHR-name-01540  
  If `name` is not `NULL`, it **must** obey any requirements listed for
  `handleType` in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#external-fence-handle-types-compatibility"
  target="_blank" rel="noopener">external fence handle types
  compatibility</a>

Valid Usage (Implicit)

- <a href="#VUID-VkImportFenceWin32HandleInfoKHR-sType-sType"
  id="VUID-VkImportFenceWin32HandleInfoKHR-sType-sType"></a>
  VUID-VkImportFenceWin32HandleInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMPORT_FENCE_WIN32_HANDLE_INFO_KHR`

- <a href="#VUID-VkImportFenceWin32HandleInfoKHR-pNext-pNext"
  id="VUID-VkImportFenceWin32HandleInfoKHR-pNext-pNext"></a>
  VUID-VkImportFenceWin32HandleInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkImportFenceWin32HandleInfoKHR-fence-parameter"
  id="VUID-VkImportFenceWin32HandleInfoKHR-fence-parameter"></a>
  VUID-VkImportFenceWin32HandleInfoKHR-fence-parameter  
  `fence` **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html) handle

- <a href="#VUID-VkImportFenceWin32HandleInfoKHR-flags-parameter"
  id="VUID-VkImportFenceWin32HandleInfoKHR-flags-parameter"></a>
  VUID-VkImportFenceWin32HandleInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of
  [VkFenceImportFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceImportFlagBits.html) values

Host Synchronization

- Host access to `fence` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_fence_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_fence_win32.html),
[VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html),
[VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html), [VkFenceImportFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceImportFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkImportFenceWin32HandleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkImportFenceWin32HandleKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImportFenceWin32HandleInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
