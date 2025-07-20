# vkCopyMemoryToAccelerationStructureKHR(3) Manual Page

## Name

vkCopyMemoryToAccelerationStructureKHR - Deserialize an acceleration structure on the host



## [](#_c_specification)C Specification

To copy host accessible memory to an acceleration structure, call:

```c++
// Provided by VK_KHR_acceleration_structure
VkResult vkCopyMemoryToAccelerationStructureKHR(
    VkDevice                                    device,
    VkDeferredOperationKHR                      deferredOperation,
    const VkCopyMemoryToAccelerationStructureInfoKHR* pInfo);
```

## [](#_parameters)Parameters

- `device` is the device which owns `pInfo->dst`.
- `deferredOperation` is an optional [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) to [request deferral](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#deferred-host-operations-requesting) for this command.
- `pInfo` is a pointer to a [VkCopyMemoryToAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToAccelerationStructureInfoKHR.html) structure defining the copy operation.

## [](#_description)Description

This command fulfills the same task as [vkCmdCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryToAccelerationStructureKHR.html) but is executed by the host.

This command can accept acceleration structures produced by either [vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html) or [vkCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyAccelerationStructureToMemoryKHR.html).

Valid Usage

- [](#VUID-vkCopyMemoryToAccelerationStructureKHR-accelerationStructureHostCommands-03583)VUID-vkCopyMemoryToAccelerationStructureKHR-accelerationStructureHostCommands-03583  
  The [`VkPhysicalDeviceAccelerationStructureFeaturesKHR`::`accelerationStructureHostCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-accelerationStructureHostCommands) feature **must** be enabled

<!--THE END-->

- [](#VUID-vkCopyMemoryToAccelerationStructureKHR-deferredOperation-03678)VUID-vkCopyMemoryToAccelerationStructureKHR-deferredOperation-03678  
  Any previous deferred operation that was associated with `deferredOperation` **must** be complete
- [](#VUID-vkCopyMemoryToAccelerationStructureKHR-pInfo-03729)VUID-vkCopyMemoryToAccelerationStructureKHR-pInfo-03729  
  `pInfo->src.hostAddress` **must** be a valid host pointer
- [](#VUID-vkCopyMemoryToAccelerationStructureKHR-pInfo-03750)VUID-vkCopyMemoryToAccelerationStructureKHR-pInfo-03750  
  `pInfo->src.hostAddress` **must** be aligned to 16 bytes
- [](#VUID-vkCopyMemoryToAccelerationStructureKHR-buffer-03730)VUID-vkCopyMemoryToAccelerationStructureKHR-buffer-03730  
  The `buffer` used to create `pInfo->dst` **must** be bound to host-visible device memory
- [](#VUID-vkCopyMemoryToAccelerationStructureKHR-buffer-03782)VUID-vkCopyMemoryToAccelerationStructureKHR-buffer-03782  
  The `buffer` used to create `pInfo->dst` **must** be bound to memory that was not allocated with multiple instances

Valid Usage (Implicit)

- [](#VUID-vkCopyMemoryToAccelerationStructureKHR-device-parameter)VUID-vkCopyMemoryToAccelerationStructureKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCopyMemoryToAccelerationStructureKHR-deferredOperation-parameter)VUID-vkCopyMemoryToAccelerationStructureKHR-deferredOperation-parameter  
  If `deferredOperation` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `deferredOperation` **must** be a valid [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) handle
- [](#VUID-vkCopyMemoryToAccelerationStructureKHR-pInfo-parameter)VUID-vkCopyMemoryToAccelerationStructureKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkCopyMemoryToAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToAccelerationStructureInfoKHR.html) structure
- [](#VUID-vkCopyMemoryToAccelerationStructureKHR-deferredOperation-parent)VUID-vkCopyMemoryToAccelerationStructureKHR-deferredOperation-parent  
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

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkCopyMemoryToAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToAccelerationStructureInfoKHR.html), [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCopyMemoryToAccelerationStructureKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0