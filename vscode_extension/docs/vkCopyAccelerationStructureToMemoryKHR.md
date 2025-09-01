# vkCopyAccelerationStructureToMemoryKHR(3) Manual Page

## Name

vkCopyAccelerationStructureToMemoryKHR - Serialize an acceleration structure on the host



## [](#_c_specification)C Specification

To copy an acceleration structure to host accessible memory, call:

```c++
// Provided by VK_KHR_acceleration_structure
VkResult vkCopyAccelerationStructureToMemoryKHR(
    VkDevice                                    device,
    VkDeferredOperationKHR                      deferredOperation,
    const VkCopyAccelerationStructureToMemoryInfoKHR* pInfo);
```

## [](#_parameters)Parameters

- `device` is the device which owns `pInfo->src`.
- `deferredOperation` is an optional [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) to [request deferral](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#deferred-host-operations-requesting) for this command.
- `pInfo` is a pointer to a [VkCopyAccelerationStructureToMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureToMemoryInfoKHR.html) structure defining the copy operation.

## [](#_description)Description

This command fulfills the same task as [vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html) but is executed by the host.

This command produces the same results as [vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html), but writes its result directly to a host pointer, and is executed on the host rather than the device. The output **may** not necessarily be bit-for-bit identical, but it can be equally used by either [vkCmdCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryToAccelerationStructureKHR.html) or [vkCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMemoryToAccelerationStructureKHR.html).

Valid Usage

- [](#VUID-vkCopyAccelerationStructureToMemoryKHR-accelerationStructureHostCommands-03584)VUID-vkCopyAccelerationStructureToMemoryKHR-accelerationStructureHostCommands-03584  
  The [`VkPhysicalDeviceAccelerationStructureFeaturesKHR`::`accelerationStructureHostCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-accelerationStructureHostCommands) feature **must** be enabled

<!--THE END-->

- [](#VUID-vkCopyAccelerationStructureToMemoryKHR-deferredOperation-03678)VUID-vkCopyAccelerationStructureToMemoryKHR-deferredOperation-03678  
  Any previous deferred operation that was associated with `deferredOperation` **must** be complete
- [](#VUID-vkCopyAccelerationStructureToMemoryKHR-buffer-03731)VUID-vkCopyAccelerationStructureToMemoryKHR-buffer-03731  
  The `buffer` used to create `pInfo->src` **must** be bound to host-visible device memory
- [](#VUID-vkCopyAccelerationStructureToMemoryKHR-pInfo-03732)VUID-vkCopyAccelerationStructureToMemoryKHR-pInfo-03732  
  `pInfo->dst.hostAddress` **must** be a valid host pointer
- [](#VUID-vkCopyAccelerationStructureToMemoryKHR-pInfo-03751)VUID-vkCopyAccelerationStructureToMemoryKHR-pInfo-03751  
  `pInfo->dst.hostAddress` **must** be aligned to 16 bytes
- [](#VUID-vkCopyAccelerationStructureToMemoryKHR-buffer-03783)VUID-vkCopyAccelerationStructureToMemoryKHR-buffer-03783  
  The `buffer` used to create `pInfo->src` **must** be bound to memory that was not allocated with multiple instances

Valid Usage (Implicit)

- [](#VUID-vkCopyAccelerationStructureToMemoryKHR-device-parameter)VUID-vkCopyAccelerationStructureToMemoryKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCopyAccelerationStructureToMemoryKHR-deferredOperation-parameter)VUID-vkCopyAccelerationStructureToMemoryKHR-deferredOperation-parameter  
  If `deferredOperation` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `deferredOperation` **must** be a valid [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) handle
- [](#VUID-vkCopyAccelerationStructureToMemoryKHR-pInfo-parameter)VUID-vkCopyAccelerationStructureToMemoryKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkCopyAccelerationStructureToMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureToMemoryInfoKHR.html) structure
- [](#VUID-vkCopyAccelerationStructureToMemoryKHR-deferredOperation-parent)VUID-vkCopyAccelerationStructureToMemoryKHR-deferredOperation-parent  
  If `deferredOperation` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_OPERATION_DEFERRED_KHR`
- `VK_OPERATION_NOT_DEFERRED_KHR`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkCopyAccelerationStructureToMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureToMemoryInfoKHR.html), [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCopyAccelerationStructureToMemoryKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0