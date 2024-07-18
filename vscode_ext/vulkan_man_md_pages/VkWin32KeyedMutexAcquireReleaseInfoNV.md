# VkWin32KeyedMutexAcquireReleaseInfoNV(3) Manual Page

## Name

VkWin32KeyedMutexAcquireReleaseInfoNV - Use Windows keyex mutex
mechanism to synchronize work



## <a href="#_c_specification" class="anchor"></a>C Specification

When submitting work that operates on memory imported from a Direct3D 11
resource to a queue, the keyed mutex mechanism **may** be used in
addition to Vulkan semaphores to synchronize the work. Keyed mutexes are
a property of a properly created shareable Direct3D 11 resource. They
**can** only be used if the imported resource was created with the
`D3D11_RESOURCE_MISC_SHARED_KEYEDMUTEX` flag.

To acquire keyed mutexes before submitted work and/or release them
after, add a
[VkWin32KeyedMutexAcquireReleaseInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWin32KeyedMutexAcquireReleaseInfoNV.html)
structure to the `pNext` chain of the [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html)
structure.

The `VkWin32KeyedMutexAcquireReleaseInfoNV` structure is defined as:

``` c
// Provided by VK_NV_win32_keyed_mutex
typedef struct VkWin32KeyedMutexAcquireReleaseInfoNV {
    VkStructureType          sType;
    const void*              pNext;
    uint32_t                 acquireCount;
    const VkDeviceMemory*    pAcquireSyncs;
    const uint64_t*          pAcquireKeys;
    const uint32_t*          pAcquireTimeoutMilliseconds;
    uint32_t                 releaseCount;
    const VkDeviceMemory*    pReleaseSyncs;
    const uint64_t*          pReleaseKeys;
} VkWin32KeyedMutexAcquireReleaseInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `acquireCount` is the number of entries in the `pAcquireSyncs`,
  `pAcquireKeys`, and `pAcquireTimeoutMilliseconds` arrays.

- `pAcquireSyncs` is a pointer to an array of
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) objects which were imported from
  Direct3D 11 resources.

- `pAcquireKeys` is a pointer to an array of mutex key values to wait
  for prior to beginning the submitted work. Entries refer to the keyed
  mutex associated with the corresponding entries in `pAcquireSyncs`.

- `pAcquireTimeoutMilliseconds` is a pointer to an array of timeout
  values, in millisecond units, for each acquire specified in
  `pAcquireKeys`.

- `releaseCount` is the number of entries in the `pReleaseSyncs` and
  `pReleaseKeys` arrays.

- `pReleaseSyncs` is a pointer to an array of
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) objects which were imported from
  Direct3D 11 resources.

- `pReleaseKeys` is a pointer to an array of mutex key values to set
  when the submitted work has completed. Entries refer to the keyed
  mutex associated with the corresponding entries in `pReleaseSyncs`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-sType-sType"
  id="VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-sType-sType"></a>
  VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_NV`

- <a
  href="#VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-pAcquireSyncs-parameter"
  id="VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-pAcquireSyncs-parameter"></a>
  VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-pAcquireSyncs-parameter  
  If `acquireCount` is not `0`, `pAcquireSyncs` **must** be a valid
  pointer to an array of `acquireCount` valid
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) handles

- <a
  href="#VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-pAcquireKeys-parameter"
  id="VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-pAcquireKeys-parameter"></a>
  VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-pAcquireKeys-parameter  
  If `acquireCount` is not `0`, `pAcquireKeys` **must** be a valid
  pointer to an array of `acquireCount` `uint64_t` values

- <a
  href="#VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-pAcquireTimeoutMilliseconds-parameter"
  id="VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-pAcquireTimeoutMilliseconds-parameter"></a>
  VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-pAcquireTimeoutMilliseconds-parameter  
  If `acquireCount` is not `0`, `pAcquireTimeoutMilliseconds` **must**
  be a valid pointer to an array of `acquireCount` `uint32_t` values

- <a
  href="#VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-pReleaseSyncs-parameter"
  id="VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-pReleaseSyncs-parameter"></a>
  VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-pReleaseSyncs-parameter  
  If `releaseCount` is not `0`, `pReleaseSyncs` **must** be a valid
  pointer to an array of `releaseCount` valid
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) handles

- <a
  href="#VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-pReleaseKeys-parameter"
  id="VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-pReleaseKeys-parameter"></a>
  VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-pReleaseKeys-parameter  
  If `releaseCount` is not `0`, `pReleaseKeys` **must** be a valid
  pointer to an array of `releaseCount` `uint64_t` values

- <a href="#VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-commonparent"
  id="VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-commonparent"></a>
  VUID-VkWin32KeyedMutexAcquireReleaseInfoNV-commonparent  
  Both of the elements of `pAcquireSyncs`, and the elements of
  `pReleaseSyncs` that are valid handles of non-ignored parameters
  **must** have been created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_win32_keyed_mutex](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_win32_keyed_mutex.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkWin32KeyedMutexAcquireReleaseInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
