# vkCopyAccelerationStructureKHR(3) Manual Page

## Name

vkCopyAccelerationStructureKHR - Copy an acceleration structure on the host



## [](#_c_specification)C Specification

To copy or compact an acceleration structure on the host, call:

```c++
// Provided by VK_KHR_acceleration_structure
VkResult vkCopyAccelerationStructureKHR(
    VkDevice                                    device,
    VkDeferredOperationKHR                      deferredOperation,
    const VkCopyAccelerationStructureInfoKHR*   pInfo);
```

## [](#_parameters)Parameters

- `device` is the device which owns the acceleration structures.
- `deferredOperation` is an optional [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) to [request deferral](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#deferred-host-operations-requesting) for this command.
- `pInfo` is a pointer to a [VkCopyAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureInfoKHR.html) structure defining the copy operation.

## [](#_description)Description

This command fulfills the same task as [vkCmdCopyAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyAccelerationStructureKHR.html) but is executed by the host.

Valid Usage

- [](#VUID-vkCopyAccelerationStructureKHR-accelerationStructureHostCommands-03582)VUID-vkCopyAccelerationStructureKHR-accelerationStructureHostCommands-03582  
  The [`VkPhysicalDeviceAccelerationStructureFeaturesKHR`::`accelerationStructureHostCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-accelerationStructureHostCommands) feature **must** be enabled

<!--THE END-->

- [](#VUID-vkCopyAccelerationStructureKHR-deferredOperation-03678)VUID-vkCopyAccelerationStructureKHR-deferredOperation-03678  
  Any previous deferred operation that was associated with `deferredOperation` **must** be complete
- [](#VUID-vkCopyAccelerationStructureKHR-buffer-03727)VUID-vkCopyAccelerationStructureKHR-buffer-03727  
  The `buffer` used to create `pInfo->src` **must** be bound to host-visible device memory
- [](#VUID-vkCopyAccelerationStructureKHR-buffer-03728)VUID-vkCopyAccelerationStructureKHR-buffer-03728  
  The `buffer` used to create `pInfo->dst` **must** be bound to host-visible device memory
- [](#VUID-vkCopyAccelerationStructureKHR-buffer-03780)VUID-vkCopyAccelerationStructureKHR-buffer-03780  
  The `buffer` used to create `pInfo->src` **must** be bound to memory that was not allocated with multiple instances
- [](#VUID-vkCopyAccelerationStructureKHR-buffer-03781)VUID-vkCopyAccelerationStructureKHR-buffer-03781  
  The `buffer` used to create `pInfo->dst` **must** be bound to memory that was not allocated with multiple instances

Valid Usage (Implicit)

- [](#VUID-vkCopyAccelerationStructureKHR-device-parameter)VUID-vkCopyAccelerationStructureKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCopyAccelerationStructureKHR-deferredOperation-parameter)VUID-vkCopyAccelerationStructureKHR-deferredOperation-parameter  
  If `deferredOperation` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `deferredOperation` **must** be a valid [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) handle
- [](#VUID-vkCopyAccelerationStructureKHR-pInfo-parameter)VUID-vkCopyAccelerationStructureKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkCopyAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureInfoKHR.html) structure
- [](#VUID-vkCopyAccelerationStructureKHR-deferredOperation-parent)VUID-vkCopyAccelerationStructureKHR-deferredOperation-parent  
  If `deferredOperation` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_OPERATION_DEFERRED_KHR`
- `VK_OPERATION_NOT_DEFERRED_KHR`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkCopyAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureInfoKHR.html), [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCopyAccelerationStructureKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0