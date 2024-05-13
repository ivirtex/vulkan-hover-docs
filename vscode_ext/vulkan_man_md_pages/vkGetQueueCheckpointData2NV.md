# vkGetQueueCheckpointData2NV(3) Manual Page

## Name

vkGetQueueCheckpointData2NV - Retrieve diagnostic checkpoint data



## <a href="#_c_specification" class="anchor"></a>C Specification

If the device encounters an error during execution, the implementation
will return a `VK_ERROR_DEVICE_LOST` error to the application at some
point during host execution. When this happens, the application **can**
call [vkGetQueueCheckpointData2NV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetQueueCheckpointData2NV.html) to
retrieve information on the most recent diagnostic checkpoints that were
executed by the device.

``` c
// Provided by VK_KHR_synchronization2 with VK_NV_device_diagnostic_checkpoints
void vkGetQueueCheckpointData2NV(
    VkQueue                                     queue,
    uint32_t*                                   pCheckpointDataCount,
    VkCheckpointData2NV*                        pCheckpointData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `queue` is the [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) object the caller would like to
  retrieve checkpoint data for

- `pCheckpointDataCount` is a pointer to an integer related to the
  number of checkpoint markers available or queried, as described below.

- `pCheckpointData` is either `NULL` or a pointer to an array of
  `VkCheckpointData2NV` structures.

## <a href="#_description" class="anchor"></a>Description

If `pCheckpointData` is `NULL`, then the number of checkpoint markers
available is returned in `pCheckpointDataCount`. Otherwise,
`pCheckpointDataCount` **must** point to a variable set by the user to
the number of elements in the `pCheckpointData` array, and on return the
variable is overwritten with the number of structures actually written
to `pCheckpointData`.

If `pCheckpointDataCount` is less than the number of checkpoint markers
available, at most `pCheckpointDataCount` structures will be written.

Valid Usage

- <a href="#VUID-vkGetQueueCheckpointData2NV-queue-03892"
  id="VUID-vkGetQueueCheckpointData2NV-queue-03892"></a>
  VUID-vkGetQueueCheckpointData2NV-queue-03892  
  The device that `queue` belongs to **must** be in the lost state

Valid Usage (Implicit)

- <a href="#VUID-vkGetQueueCheckpointData2NV-queue-parameter"
  id="VUID-vkGetQueueCheckpointData2NV-queue-parameter"></a>
  VUID-vkGetQueueCheckpointData2NV-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) handle

- <a
  href="#VUID-vkGetQueueCheckpointData2NV-pCheckpointDataCount-parameter"
  id="VUID-vkGetQueueCheckpointData2NV-pCheckpointDataCount-parameter"></a>
  VUID-vkGetQueueCheckpointData2NV-pCheckpointDataCount-parameter  
  `pCheckpointDataCount` **must** be a valid pointer to a `uint32_t`
  value

- <a href="#VUID-vkGetQueueCheckpointData2NV-pCheckpointData-parameter"
  id="VUID-vkGetQueueCheckpointData2NV-pCheckpointData-parameter"></a>
  VUID-vkGetQueueCheckpointData2NV-pCheckpointData-parameter  
  If the value referenced by `pCheckpointDataCount` is not `0`, and
  `pCheckpointData` is not `NULL`, `pCheckpointData` **must** be a valid
  pointer to an array of `pCheckpointDataCount`
  [VkCheckpointData2NV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCheckpointData2NV.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html),
[VK_NV_device_diagnostic_checkpoints](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_diagnostic_checkpoints.html),
[VkCheckpointData2NV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCheckpointData2NV.html), [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetQueueCheckpointData2NV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
