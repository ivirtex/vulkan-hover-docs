# VkExternalFenceHandleTypeFlagBits(3) Manual Page

## Name

VkExternalFenceHandleTypeFlagBits - Bitmask of valid external fence
handle types



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in

- [VkPhysicalDeviceExternalFenceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalFenceInfo.html)::`handleType`

- [VkExternalFenceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceProperties.html)::`exportFromImportedHandleTypes`

- [VkExternalFenceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceProperties.html)::`compatibleHandleTypes`

indicate external fence handle types, and are:

``` c
// Provided by VK_VERSION_1_1
typedef enum VkExternalFenceHandleTypeFlagBits {
    VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT = 0x00000001,
    VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT = 0x00000002,
    VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT = 0x00000004,
    VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT = 0x00000008,
  // Provided by VK_KHR_external_fence_capabilities
    VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT_KHR = VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT,
  // Provided by VK_KHR_external_fence_capabilities
    VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR = VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT,
  // Provided by VK_KHR_external_fence_capabilities
    VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR = VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT,
  // Provided by VK_KHR_external_fence_capabilities
    VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT_KHR = VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT,
} VkExternalFenceHandleTypeFlagBits;
```

or the equivalent

``` c
// Provided by VK_KHR_external_fence_capabilities
typedef VkExternalFenceHandleTypeFlagBits VkExternalFenceHandleTypeFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT` specifies a POSIX file
  descriptor handle that has only limited valid usage outside of Vulkan
  and other compatible APIs. It **must** be compatible with the POSIX
  system calls `dup`, `dup2`, `close`, and the non-standard system call
  `dup3`. Additionally, it **must** be transportable over a socket using
  an `SCM_RIGHTS` control message. It owns a reference to the underlying
  synchronization primitive represented by its Vulkan fence object.

- `VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT` specifies an NT
  handle that has only limited valid usage outside of Vulkan and other
  compatible APIs. It **must** be compatible with the functions
  `DuplicateHandle`, `CloseHandle`, `CompareObjectHandles`,
  `GetHandleInformation`, and `SetHandleInformation`. It owns a
  reference to the underlying synchronization primitive represented by
  its Vulkan fence object.

- `VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT` specifies a
  global share handle that has only limited valid usage outside of
  Vulkan and other compatible APIs. It is not compatible with any native
  APIs. It does not own a reference to the underlying synchronization
  primitive represented by its Vulkan fence object, and will therefore
  become invalid when all Vulkan fence objects associated with it are
  destroyed.

- `VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT` specifies a POSIX file
  descriptor handle to a Linux Sync File or Android Fence. It can be
  used with any native API accepting a valid sync file or fence as
  input. It owns a reference to the underlying synchronization primitive
  associated with the file descriptor. Implementations which support
  importing this handle type **must** accept any type of sync or fence
  FD supported by the native system they are running on.

Some external fence handle types can only be shared within the same
underlying physical device and/or the same driver version, as defined in
the following table:

|  |  |  |
|----|----|----|
| Handle type | `VkPhysicalDeviceIDProperties`::`driverUUID` | `VkPhysicalDeviceIDProperties`::`deviceUUID` |
| `VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT` | Must match | Must match |
| `VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT` | Must match | Must match |
| `VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT` | Must match | Must match |
| `VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT` | No restriction | No restriction |

Table 1. External fence handle types compatibility
{#external-fence-handle-types-compatibility}

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkExternalFenceHandleTypeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlags.html),
[VkFenceGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceGetFdInfoKHR.html),
[VkFenceGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceGetWin32HandleInfoKHR.html),
[VkImportFenceFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportFenceFdInfoKHR.html),
[VkImportFenceWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportFenceWin32HandleInfoKHR.html),
[VkPhysicalDeviceExternalFenceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalFenceInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExternalFenceHandleTypeFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
