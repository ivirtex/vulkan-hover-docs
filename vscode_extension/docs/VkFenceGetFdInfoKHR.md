# VkFenceGetFdInfoKHR(3) Manual Page

## Name

VkFenceGetFdInfoKHR - Structure describing a POSIX FD fence export operation



## [](#_c_specification)C Specification

The `VkFenceGetFdInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_external_fence_fd
typedef struct VkFenceGetFdInfoKHR {
    VkStructureType                      sType;
    const void*                          pNext;
    VkFence                              fence;
    VkExternalFenceHandleTypeFlagBits    handleType;
} VkFenceGetFdInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `fence` is the fence from which state will be exported.
- `handleType` is a [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlagBits.html) value specifying the type of handle requested.

## [](#_description)Description

The properties of the file descriptor returned depend on the value of `handleType`. See [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlagBits.html) for a description of the properties of the defined external fence handle types.

Valid Usage

- [](#VUID-VkFenceGetFdInfoKHR-handleType-01453)VUID-VkFenceGetFdInfoKHR-handleType-01453  
  `handleType` **must** have been included in [VkExportFenceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportFenceCreateInfo.html)::`handleTypes` when `fence`’s current payload was created
- [](#VUID-VkFenceGetFdInfoKHR-handleType-01454)VUID-VkFenceGetFdInfoKHR-handleType-01454  
  If `handleType` refers to a handle type with copy payload transference semantics, `fence` **must** be signaled, or have an associated [fence signal operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-fences-signaling) pending execution
- [](#VUID-VkFenceGetFdInfoKHR-fence-01455)VUID-VkFenceGetFdInfoKHR-fence-01455  
  `fence` **must** not currently have its payload replaced by an imported payload as described below in [Importing Fence Payloads](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-fences-importing) unless that imported payload’s handle type was included in [VkExternalFenceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceProperties.html)::`exportFromImportedHandleTypes` for `handleType`
- [](#VUID-VkFenceGetFdInfoKHR-handleType-01456)VUID-VkFenceGetFdInfoKHR-handleType-01456  
  `handleType` **must** be defined as a POSIX file descriptor handle

Valid Usage (Implicit)

- [](#VUID-VkFenceGetFdInfoKHR-sType-sType)VUID-VkFenceGetFdInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FENCE_GET_FD_INFO_KHR`
- [](#VUID-VkFenceGetFdInfoKHR-pNext-pNext)VUID-VkFenceGetFdInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkFenceGetFdInfoKHR-fence-parameter)VUID-VkFenceGetFdInfoKHR-fence-parameter  
  `fence` **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html) handle
- [](#VUID-VkFenceGetFdInfoKHR-handleType-parameter)VUID-VkFenceGetFdInfoKHR-handleType-parameter  
  `handleType` **must** be a valid [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlagBits.html) value

## [](#_see_also)See Also

[VK\_KHR\_external\_fence\_fd](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_fence_fd.html), [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlagBits.html), [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetFenceFdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetFenceFdKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFenceGetFdInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0