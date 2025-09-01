# VkFenceGetWin32HandleInfoKHR(3) Manual Page

## Name

VkFenceGetWin32HandleInfoKHR - Structure describing a Win32 handle fence export operation



## [](#_c_specification)C Specification

The `VkFenceGetWin32HandleInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_external_fence_win32
typedef struct VkFenceGetWin32HandleInfoKHR {
    VkStructureType                      sType;
    const void*                          pNext;
    VkFence                              fence;
    VkExternalFenceHandleTypeFlagBits    handleType;
} VkFenceGetWin32HandleInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `fence` is the fence from which state will be exported.
- `handleType` is a [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlagBits.html) value specifying the type of handle requested.

## [](#_description)Description

The properties of the handle returned depend on the value of `handleType`. See [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlagBits.html) for a description of the properties of the defined external fence handle types.

Valid Usage

- [](#VUID-VkFenceGetWin32HandleInfoKHR-handleType-01448)VUID-VkFenceGetWin32HandleInfoKHR-handleType-01448  
  `handleType` **must** have been included in [VkExportFenceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportFenceCreateInfo.html)::`handleTypes` when the `fence`’s current payload was created
- [](#VUID-VkFenceGetWin32HandleInfoKHR-handleType-01449)VUID-VkFenceGetWin32HandleInfoKHR-handleType-01449  
  If `handleType` is defined as an NT handle, [vkGetFenceWin32HandleKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetFenceWin32HandleKHR.html) **must** be called no more than once for each valid unique combination of `fence` and `handleType`
- [](#VUID-VkFenceGetWin32HandleInfoKHR-fence-01450)VUID-VkFenceGetWin32HandleInfoKHR-fence-01450  
  `fence` **must** not currently have its payload replaced by an imported payload as described below in [Importing Fence Payloads](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-fences-importing) unless that imported payload’s handle type was included in [VkExternalFenceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceProperties.html)::`exportFromImportedHandleTypes` for `handleType`
- [](#VUID-VkFenceGetWin32HandleInfoKHR-handleType-01451)VUID-VkFenceGetWin32HandleInfoKHR-handleType-01451  
  If `handleType` refers to a handle type with copy payload transference semantics, `fence` **must** be signaled, or have an associated [fence signal operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-fences-signaling) pending execution
- [](#VUID-VkFenceGetWin32HandleInfoKHR-handleType-01452)VUID-VkFenceGetWin32HandleInfoKHR-handleType-01452  
  `handleType` **must** be defined as an NT handle or a global share handle

Valid Usage (Implicit)

- [](#VUID-VkFenceGetWin32HandleInfoKHR-sType-sType)VUID-VkFenceGetWin32HandleInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FENCE_GET_WIN32_HANDLE_INFO_KHR`
- [](#VUID-VkFenceGetWin32HandleInfoKHR-pNext-pNext)VUID-VkFenceGetWin32HandleInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkFenceGetWin32HandleInfoKHR-fence-parameter)VUID-VkFenceGetWin32HandleInfoKHR-fence-parameter  
  `fence` **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html) handle
- [](#VUID-VkFenceGetWin32HandleInfoKHR-handleType-parameter)VUID-VkFenceGetWin32HandleInfoKHR-handleType-parameter  
  `handleType` **must** be a valid [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlagBits.html) value

## [](#_see_also)See Also

[VK\_KHR\_external\_fence\_win32](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_fence_win32.html), [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlagBits.html), [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetFenceWin32HandleKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetFenceWin32HandleKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFenceGetWin32HandleInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0