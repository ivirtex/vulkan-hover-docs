# vkGetQueueCheckpointDataNV(3) Manual Page

## Name

vkGetQueueCheckpointDataNV - Retrieve diagnostic checkpoint data



## [](#_c_specification)C Specification

If the device encounters an error during execution, the implementation will return a `VK_ERROR_DEVICE_LOST` error to the application at a certain point during host execution. When this happens, the application **can** call [vkGetQueueCheckpointDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetQueueCheckpointDataNV.html) to retrieve information on the most recent diagnostic checkpoints that were executed by the device.

```c++
// Provided by VK_NV_device_diagnostic_checkpoints
void vkGetQueueCheckpointDataNV(
    VkQueue                                     queue,
    uint32_t*                                   pCheckpointDataCount,
    VkCheckpointDataNV*                         pCheckpointData);
```

## [](#_parameters)Parameters

- `queue` is the [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) object the caller would like to retrieve checkpoint data for
- `pCheckpointDataCount` is a pointer to an integer related to the number of checkpoint markers available or queried, as described below.
- `pCheckpointData` is either `NULL` or a pointer to an array of `VkCheckpointDataNV` structures.

## [](#_description)Description

If `pCheckpointData` is `NULL`, then the number of checkpoint markers available is returned in `pCheckpointDataCount`.

Otherwise, `pCheckpointDataCount` **must** point to a variable set by the application to the number of elements in the `pCheckpointData` array, and on return the variable is overwritten with the number of structures actually written to `pCheckpointData`.

If `pCheckpointDataCount` is less than the number of checkpoint markers available, at most `pCheckpointDataCount` structures will be written.

Valid Usage

- [](#VUID-vkGetQueueCheckpointDataNV-queue-02025)VUID-vkGetQueueCheckpointDataNV-queue-02025  
  The device that `queue` belongs to **must** be in the lost state

Valid Usage (Implicit)

- [](#VUID-vkGetQueueCheckpointDataNV-queue-parameter)VUID-vkGetQueueCheckpointDataNV-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) handle
- [](#VUID-vkGetQueueCheckpointDataNV-pCheckpointDataCount-parameter)VUID-vkGetQueueCheckpointDataNV-pCheckpointDataCount-parameter  
  `pCheckpointDataCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetQueueCheckpointDataNV-pCheckpointData-parameter)VUID-vkGetQueueCheckpointDataNV-pCheckpointData-parameter  
  If the value referenced by `pCheckpointDataCount` is not `0`, and `pCheckpointData` is not `NULL`, `pCheckpointData` **must** be a valid pointer to an array of `pCheckpointDataCount` [VkCheckpointDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCheckpointDataNV.html) structures

## [](#_see_also)See Also

[VK\_NV\_device\_diagnostic\_checkpoints](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_diagnostic_checkpoints.html), [VkCheckpointDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCheckpointDataNV.html), [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetQueueCheckpointDataNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0