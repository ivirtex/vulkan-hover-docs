# VkWin32KeyedMutexAcquireReleaseInfoKHR(3) Manual Page

## Name

VkWin32KeyedMutexAcquireReleaseInfoKHR - Use the Windows keyed mutex mechanism to synchronize work



## [](#_c_specification)C Specification

When submitting work that operates on memory imported from a Direct3D 11 resource to a queue, the keyed mutex mechanism **may** be used in addition to Vulkan semaphores to synchronize the work. Keyed mutexes are a property of a properly created shareable Direct3D 11 resource. They **can** only be used if the imported resource was created with the `D3D11_RESOURCE_MISC_SHARED_KEYEDMUTEX` flag.

To acquire keyed mutexes before submitted work and/or release them after, add a [VkWin32KeyedMutexAcquireReleaseInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWin32KeyedMutexAcquireReleaseInfoKHR.html) structure to the `pNext` chain of the [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html) structure.

The `VkWin32KeyedMutexAcquireReleaseInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_win32_keyed_mutex
typedef struct VkWin32KeyedMutexAcquireReleaseInfoKHR {
    VkStructureType          sType;
    const void*              pNext;
    uint32_t                 acquireCount;
    const VkDeviceMemory*    pAcquireSyncs;
    const uint64_t*          pAcquireKeys;
    const uint32_t*          pAcquireTimeouts;
    uint32_t                 releaseCount;
    const VkDeviceMemory*    pReleaseSyncs;
    const uint64_t*          pReleaseKeys;
} VkWin32KeyedMutexAcquireReleaseInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `acquireCount` is the number of entries in the `pAcquireSyncs`, `pAcquireKeys`, and `pAcquireTimeouts` arrays.
- `pAcquireSyncs` is a pointer to an array of [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) objects which were imported from Direct3D 11 resources.
- `pAcquireKeys` is a pointer to an array of mutex key values to wait for prior to beginning the submitted work. Entries refer to the keyed mutex associated with the corresponding entries in `pAcquireSyncs`.
- `pAcquireTimeouts` is a pointer to an array of timeout values, in millisecond units, for each acquire specified in `pAcquireKeys`.
- `releaseCount` is the number of entries in the `pReleaseSyncs` and `pReleaseKeys` arrays.
- `pReleaseSyncs` is a pointer to an array of [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) objects which were imported from Direct3D 11 resources.
- `pReleaseKeys` is a pointer to an array of mutex key values to set when the submitted work has completed. Entries refer to the keyed mutex associated with the corresponding entries in `pReleaseSyncs`.

## [](#_description)Description

Valid Usage

- [](#VUID-VkWin32KeyedMutexAcquireReleaseInfoKHR-pAcquireSyncs-00081)VUID-VkWin32KeyedMutexAcquireReleaseInfoKHR-pAcquireSyncs-00081  
  Each member of `pAcquireSyncs` and `pReleaseSyncs` **must** be a device memory object imported by setting [VkImportMemoryWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMemoryWin32HandleInfoKHR.html)::`handleType` to `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT` or `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT`

Valid Usage (Implicit)

- [](#VUID-VkWin32KeyedMutexAcquireReleaseInfoKHR-sType-sType)VUID-VkWin32KeyedMutexAcquireReleaseInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_KHR`
- [](#VUID-VkWin32KeyedMutexAcquireReleaseInfoKHR-pAcquireSyncs-parameter)VUID-VkWin32KeyedMutexAcquireReleaseInfoKHR-pAcquireSyncs-parameter  
  If `acquireCount` is not `0`, `pAcquireSyncs` **must** be a valid pointer to an array of `acquireCount` valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handles
- [](#VUID-VkWin32KeyedMutexAcquireReleaseInfoKHR-pAcquireKeys-parameter)VUID-VkWin32KeyedMutexAcquireReleaseInfoKHR-pAcquireKeys-parameter  
  If `acquireCount` is not `0`, `pAcquireKeys` **must** be a valid pointer to an array of `acquireCount` `uint64_t` values
- [](#VUID-VkWin32KeyedMutexAcquireReleaseInfoKHR-pAcquireTimeouts-parameter)VUID-VkWin32KeyedMutexAcquireReleaseInfoKHR-pAcquireTimeouts-parameter  
  If `acquireCount` is not `0`, `pAcquireTimeouts` **must** be a valid pointer to an array of `acquireCount` `uint32_t` values
- [](#VUID-VkWin32KeyedMutexAcquireReleaseInfoKHR-pReleaseSyncs-parameter)VUID-VkWin32KeyedMutexAcquireReleaseInfoKHR-pReleaseSyncs-parameter  
  If `releaseCount` is not `0`, `pReleaseSyncs` **must** be a valid pointer to an array of `releaseCount` valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handles
- [](#VUID-VkWin32KeyedMutexAcquireReleaseInfoKHR-pReleaseKeys-parameter)VUID-VkWin32KeyedMutexAcquireReleaseInfoKHR-pReleaseKeys-parameter  
  If `releaseCount` is not `0`, `pReleaseKeys` **must** be a valid pointer to an array of `releaseCount` `uint64_t` values
- [](#VUID-VkWin32KeyedMutexAcquireReleaseInfoKHR-commonparent)VUID-VkWin32KeyedMutexAcquireReleaseInfoKHR-commonparent  
  Both of the elements of `pAcquireSyncs`, and the elements of `pReleaseSyncs` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_KHR\_win32\_keyed\_mutex](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_win32_keyed_mutex.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkWin32KeyedMutexAcquireReleaseInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0