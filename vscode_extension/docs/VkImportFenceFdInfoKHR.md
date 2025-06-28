# VkImportFenceFdInfoKHR(3) Manual Page

## Name

VkImportFenceFdInfoKHR - (None)



## [](#_c_specification)C Specification

The `VkImportFenceFdInfoKHR` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `fence` is the fence into which the payload will be imported.
- `flags` is a bitmask of [VkFenceImportFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceImportFlagBits.html) specifying additional parameters for the fence payload import operation.
- `handleType` is a [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlagBits.html) value specifying the type of `fd`.
- `fd` is the external handle to import.

## [](#_description)Description

The handle types supported by `handleType` are:

Table 1. Handle Types Supported by `VkImportFenceFdInfoKHR`    Handle Type Transference Permanence Supported

`VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT`

Reference

Temporary,Permanent

`VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT`

Copy

Temporary

Valid Usage

- [](#VUID-VkImportFenceFdInfoKHR-handleType-01464)VUID-VkImportFenceFdInfoKHR-handleType-01464  
  `handleType` **must** be a value included in the [Handle Types Supported by `VkImportFenceFdInfoKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-fence-handletypes-fd) table
- [](#VUID-VkImportFenceFdInfoKHR-fd-01541)VUID-VkImportFenceFdInfoKHR-fd-01541  
  `fd` **must** obey any requirements listed for `handleType` in [external fence handle types compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#external-fence-handle-types-compatibility)
- [](#VUID-VkImportFenceFdInfoKHR-handleType-07306)VUID-VkImportFenceFdInfoKHR-handleType-07306  
  If `handleType` refers to a handle type with copy payload transference semantics, `flags` **must** contain `VK_FENCE_IMPORT_TEMPORARY_BIT`

If `handleType` is `VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT`, the special value `-1` for `fd` is treated like a valid sync file descriptor referring to an object that has already signaled. The import operation will succeed and the `VkFence` will have a temporarily imported payload as if a valid file descriptor had been provided.

Note

This special behavior for importing an invalid sync file descriptor allows easier interoperability with other system APIs which use the convention that an invalid sync file descriptor represents work that has already completed and does not need to be waited for. It is consistent with the option for implementations to return a `-1` file descriptor when exporting a `VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT` from a `VkFence` which is signaled.

Valid Usage (Implicit)

- [](#VUID-VkImportFenceFdInfoKHR-sType-sType)VUID-VkImportFenceFdInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_FENCE_FD_INFO_KHR`
- [](#VUID-VkImportFenceFdInfoKHR-pNext-pNext)VUID-VkImportFenceFdInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkImportFenceFdInfoKHR-fence-parameter)VUID-VkImportFenceFdInfoKHR-fence-parameter  
  `fence` **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html) handle
- [](#VUID-VkImportFenceFdInfoKHR-flags-parameter)VUID-VkImportFenceFdInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of [VkFenceImportFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceImportFlagBits.html) values
- [](#VUID-VkImportFenceFdInfoKHR-handleType-parameter)VUID-VkImportFenceFdInfoKHR-handleType-parameter  
  `handleType` **must** be a valid [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlagBits.html) value

Host Synchronization

- Host access to `fence` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_KHR\_external\_fence\_fd](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_fence_fd.html), [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlagBits.html), [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html), [VkFenceImportFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceImportFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkImportFenceFdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkImportFenceFdKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImportFenceFdInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0