# VkExternalSemaphoreHandleTypeFlagBits(3) Manual Page

## Name

VkExternalSemaphoreHandleTypeFlagBits - Bitmask of valid external
semaphore handle types



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in
[VkPhysicalDeviceExternalSemaphoreInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalSemaphoreInfo.html)::`handleType`,
specifying an external semaphore handle type, are:

``` c
// Provided by VK_VERSION_1_1
typedef enum VkExternalSemaphoreHandleTypeFlagBits {
    VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT = 0x00000001,
    VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT = 0x00000002,
    VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT = 0x00000004,
    VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT = 0x00000008,
    VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT = 0x00000010,
  // Provided by VK_FUCHSIA_external_semaphore
    VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_ZIRCON_EVENT_BIT_FUCHSIA = 0x00000080,
    VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D11_FENCE_BIT = VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT,
  // Provided by VK_KHR_external_semaphore_capabilities
    VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT_KHR = VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT,
  // Provided by VK_KHR_external_semaphore_capabilities
    VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR = VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT,
  // Provided by VK_KHR_external_semaphore_capabilities
    VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR = VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT,
  // Provided by VK_KHR_external_semaphore_capabilities
    VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT_KHR = VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT,
  // Provided by VK_KHR_external_semaphore_capabilities
    VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT_KHR = VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT,
} VkExternalSemaphoreHandleTypeFlagBits;
```

or the equivalent

``` c
// Provided by VK_KHR_external_semaphore_capabilities
typedef VkExternalSemaphoreHandleTypeFlagBits VkExternalSemaphoreHandleTypeFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT` specifies a POSIX
  file descriptor handle that has only limited valid usage outside of
  Vulkan and other compatible APIs. It **must** be compatible with the
  POSIX system calls `dup`, `dup2`, `close`, and the non-standard system
  call `dup3`. Additionally, it **must** be transportable over a socket
  using an `SCM_RIGHTS` control message. It owns a reference to the
  underlying synchronization primitive represented by its Vulkan
  semaphore object.

- `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT` specifies an NT
  handle that has only limited valid usage outside of Vulkan and other
  compatible APIs. It **must** be compatible with the functions
  `DuplicateHandle`, `CloseHandle`, `CompareObjectHandles`,
  `GetHandleInformation`, and `SetHandleInformation`. It owns a
  reference to the underlying synchronization primitive represented by
  its Vulkan semaphore object.

- `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT` specifies a
  global share handle that has only limited valid usage outside of
  Vulkan and other compatible APIs. It is not compatible with any native
  APIs. It does not own a reference to the underlying synchronization
  primitive represented by its Vulkan semaphore object, and will
  therefore become invalid when all Vulkan semaphore objects associated
  with it are destroyed.

- `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT` specifies an NT
  handle returned by `ID3D12Device`::`CreateSharedHandle` referring to a
  Direct3D 12 fence, or `ID3D11Device5`::`CreateFence` referring to a
  Direct3D 11 fence. It owns a reference to the underlying
  synchronization primitive associated with the Direct3D fence.

- `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D11_FENCE_BIT` is an alias of
  `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT` with the same
  meaning. It is provided for convenience and code clarity when
  interacting with D3D11 fences.

- `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT` specifies a POSIX file
  descriptor handle to a Linux Sync File or Android Fence object. It can
  be used with any native API accepting a valid sync file or fence as
  input. It owns a reference to the underlying synchronization primitive
  associated with the file descriptor. Implementations which support
  importing this handle type **must** accept any type of sync or fence
  FD supported by the native system they are running on.

- `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_ZIRCON_EVENT_BIT_FUCHSIA` specifies
  a handle to a Zircon event object. It can be used with any native API
  that accepts a Zircon event handle. Zircon event handles are created
  with `ZX_RIGHTS_BASIC` and `ZX_RIGHTS_SIGNAL` rights. Vulkan on
  Fuchsia uses only the ZX_EVENT_SIGNALED bit when signaling or waiting.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Handles of type
<code>VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT</code> generated by
the implementation may represent either Linux Sync Files or Android
Fences at the implementation’s discretion. Applications
<strong>should</strong> only use operations defined for both types of
file descriptors, unless they know via means external to Vulkan the type
of the file descriptor, or are prepared to deal with the system-defined
operation failures resulting from using the wrong type.</p></td>
</tr>
</tbody>
</table>

Some external semaphore handle types can only be shared within the same
underlying physical device and/or the same driver version, as defined in
the following table:

|  |  |  |
|----|----|----|
| Handle type | `VkPhysicalDeviceIDProperties`::`driverUUID` | `VkPhysicalDeviceIDProperties`::`deviceUUID` |
| `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT` | Must match | Must match |
| `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT` | Must match | Must match |
| `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT` | Must match | Must match |
| `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT` | Must match | Must match |
| `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT` | No restriction | No restriction |
| `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_ZIRCON_EVENT_BIT_FUCHSIA` | No restriction | No restriction |

Table 1. External semaphore handle types compatibility
{#external-semaphore-handle-types-compatibility}

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkExternalSemaphoreHandleTypeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlags.html),
[VkImportSemaphoreFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreFdInfoKHR.html),
[VkImportSemaphoreWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreWin32HandleInfoKHR.html),
[VkImportSemaphoreZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreZirconHandleInfoFUCHSIA.html),
[VkPhysicalDeviceExternalSemaphoreInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalSemaphoreInfo.html),
[VkSemaphoreGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreGetFdInfoKHR.html),
[VkSemaphoreGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreGetWin32HandleInfoKHR.html),
[VkSemaphoreGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreGetZirconHandleInfoFUCHSIA.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExternalSemaphoreHandleTypeFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
