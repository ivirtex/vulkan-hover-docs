# VkImportFenceWin32HandleInfoKHR(3) Manual Page

## Name

VkImportFenceWin32HandleInfoKHR - (None)



## [](#_c_specification)C Specification

The `VkImportFenceWin32HandleInfoKHR` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `fence` is the fence into which the state will be imported.
- `flags` is a bitmask of [VkFenceImportFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceImportFlagBits.html) specifying additional parameters for the fence payload import operation.
- `handleType` is a [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlagBits.html) value specifying the type of `handle`.
- `handle` is `NULL` or the external handle to import.
- `name` is `NULL` or a null-terminated UTF-16 string naming the underlying synchronization primitive to import.

## [](#_description)Description

The handle types supported by `handleType` are:

Table 1. Handle Types Supported by `VkImportFenceWin32HandleInfoKHR`    Handle Type Transference Permanence Supported

`VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT`

Reference

Temporary,Permanent

`VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT`

Reference

Temporary,Permanent

Valid Usage

- [](#VUID-VkImportFenceWin32HandleInfoKHR-handleType-01457)VUID-VkImportFenceWin32HandleInfoKHR-handleType-01457  
  `handleType` **must** be a value included in the [Handle Types Supported by `VkImportFenceWin32HandleInfoKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-fence-handletypes-win32) table
- [](#VUID-VkImportFenceWin32HandleInfoKHR-handleType-01459)VUID-VkImportFenceWin32HandleInfoKHR-handleType-01459  
  If `handleType` is not `VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT`, `name` **must** be `NULL`
- [](#VUID-VkImportFenceWin32HandleInfoKHR-handleType-01460)VUID-VkImportFenceWin32HandleInfoKHR-handleType-01460  
  If `handle` is `NULL`, `name` **must** name a valid synchronization primitive of the type specified by `handleType`
- [](#VUID-VkImportFenceWin32HandleInfoKHR-handleType-01461)VUID-VkImportFenceWin32HandleInfoKHR-handleType-01461  
  If `name` is `NULL`, `handle` **must** be a valid handle of the type specified by `handleType`
- [](#VUID-VkImportFenceWin32HandleInfoKHR-handle-01462)VUID-VkImportFenceWin32HandleInfoKHR-handle-01462  
  If `handle` is not `NULL`, `name` **must** be `NULL`
- [](#VUID-VkImportFenceWin32HandleInfoKHR-handle-01539)VUID-VkImportFenceWin32HandleInfoKHR-handle-01539  
  If `handle` is not `NULL`, it **must** obey any requirements listed for `handleType` in [external fence handle types compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#external-fence-handle-types-compatibility)
- [](#VUID-VkImportFenceWin32HandleInfoKHR-name-01540)VUID-VkImportFenceWin32HandleInfoKHR-name-01540  
  If `name` is not `NULL`, it **must** obey any requirements listed for `handleType` in [external fence handle types compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#external-fence-handle-types-compatibility)

Valid Usage (Implicit)

- [](#VUID-VkImportFenceWin32HandleInfoKHR-sType-sType)VUID-VkImportFenceWin32HandleInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_FENCE_WIN32_HANDLE_INFO_KHR`
- [](#VUID-VkImportFenceWin32HandleInfoKHR-pNext-pNext)VUID-VkImportFenceWin32HandleInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkImportFenceWin32HandleInfoKHR-fence-parameter)VUID-VkImportFenceWin32HandleInfoKHR-fence-parameter  
  `fence` **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html) handle
- [](#VUID-VkImportFenceWin32HandleInfoKHR-flags-parameter)VUID-VkImportFenceWin32HandleInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of [VkFenceImportFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceImportFlagBits.html) values

Host Synchronization

- Host access to `fence` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_KHR\_external\_fence\_win32](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_fence_win32.html), [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlagBits.html), [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html), [VkFenceImportFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceImportFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkImportFenceWin32HandleKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkImportFenceWin32HandleKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImportFenceWin32HandleInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0