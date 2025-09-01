# vkGetQueueCheckpointData2NV(3) Manual Page

## Name

vkGetQueueCheckpointData2NV - Retrieve diagnostic checkpoint data



## [](#_c_specification)C Specification

If the device encounters an error during execution, the implementation will return a `VK_ERROR_DEVICE_LOST` error to the application at some point during host execution. When this happens, the application **can** call [vkGetQueueCheckpointData2NV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetQueueCheckpointData2NV.html) to retrieve information on the most recent diagnostic checkpoints that were executed by the device.

```c++
// Provided by VK_NV_device_diagnostic_checkpoints with VK_VERSION_1_3 or VK_KHR_synchronization2
void vkGetQueueCheckpointData2NV(
    VkQueue                                     queue,
    uint32_t*                                   pCheckpointDataCount,
    VkCheckpointData2NV*                        pCheckpointData);
```

## [](#_parameters)Parameters

- `queue` is the [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) object the caller would like to retrieve checkpoint data for
- `pCheckpointDataCount` is a pointer to an integer related to the number of checkpoint markers available or queried, as described below.
- `pCheckpointData` is either `NULL` or a pointer to an array of `VkCheckpointData2NV` structures.

## [](#_description)Description

If `pCheckpointData` is `NULL`, then the number of checkpoint markers available is returned in `pCheckpointDataCount`. Otherwise, `pCheckpointDataCount` **must** point to a variable set by the application to the number of elements in the `pCheckpointData` array, and on return the variable is overwritten with the number of structures actually written to `pCheckpointData`.

If `pCheckpointDataCount` is less than the number of checkpoint markers available, at most `pCheckpointDataCount` structures will be written.

Valid Usage

- [](#VUID-vkGetQueueCheckpointData2NV-queue-03892)VUID-vkGetQueueCheckpointData2NV-queue-03892  
  The device that `queue` belongs to **must** be in the lost state

Valid Usage (Implicit)

- [](#VUID-vkGetQueueCheckpointData2NV-queue-parameter)VUID-vkGetQueueCheckpointData2NV-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) handle
- [](#VUID-vkGetQueueCheckpointData2NV-pCheckpointDataCount-parameter)VUID-vkGetQueueCheckpointData2NV-pCheckpointDataCount-parameter  
  `pCheckpointDataCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetQueueCheckpointData2NV-pCheckpointData-parameter)VUID-vkGetQueueCheckpointData2NV-pCheckpointData-parameter  
  If the value referenced by `pCheckpointDataCount` is not `0`, and `pCheckpointData` is not `NULL`, `pCheckpointData` **must** be a valid pointer to an array of `pCheckpointDataCount` [VkCheckpointData2NV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCheckpointData2NV.html) structures

## [](#_see_also)See Also

[VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_NV\_device\_diagnostic\_checkpoints](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_diagnostic_checkpoints.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCheckpointData2NV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCheckpointData2NV.html), [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetQueueCheckpointData2NV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0