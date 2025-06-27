# vkGetQueueCheckpointDataNV(3) Manual Page

## Name

vkGetQueueCheckpointDataNV - Retrieve diagnostic checkpoint data



## <a href="#_c_specification" class="anchor"></a>C Specification

If the device encounters an error during execution, the implementation
will return a `VK_ERROR_DEVICE_LOST` error to the application at a
certain point during host execution. When this happens, the application
**can** call
[vkGetQueueCheckpointDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetQueueCheckpointDataNV.html) to
retrieve information on the most recent diagnostic checkpoints that were
executed by the device.

``` c
// Provided by VK_NV_device_diagnostic_checkpoints
void vkGetQueueCheckpointDataNV(
    VkQueue                                     queue,
    uint32_t*                                   pCheckpointDataCount,
    VkCheckpointDataNV*                         pCheckpointData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `queue` is the [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) object the caller would like to
  retrieve checkpoint data for

- `pCheckpointDataCount` is a pointer to an integer related to the
  number of checkpoint markers available or queried, as described below.

- `pCheckpointData` is either `NULL` or a pointer to an array of
  `VkCheckpointDataNV` structures.

## <a href="#_description" class="anchor"></a>Description

If `pCheckpointData` is `NULL`, then the number of checkpoint markers
available is returned in `pCheckpointDataCount`.

Otherwise, `pCheckpointDataCount` **must** point to a variable set by
the application to the number of elements in the `pCheckpointData`
array, and on return the variable is overwritten with the number of
structures actually written to `pCheckpointData`.

If `pCheckpointDataCount` is less than the number of checkpoint markers
available, at most `pCheckpointDataCount` structures will be written.

Valid Usage

- <a href="#VUID-vkGetQueueCheckpointDataNV-queue-02025"
  id="VUID-vkGetQueueCheckpointDataNV-queue-02025"></a>
  VUID-vkGetQueueCheckpointDataNV-queue-02025  
  The device that `queue` belongs to **must** be in the lost state

Valid Usage (Implicit)

- <a href="#VUID-vkGetQueueCheckpointDataNV-queue-parameter"
  id="VUID-vkGetQueueCheckpointDataNV-queue-parameter"></a>
  VUID-vkGetQueueCheckpointDataNV-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) handle

- <a
  href="#VUID-vkGetQueueCheckpointDataNV-pCheckpointDataCount-parameter"
  id="VUID-vkGetQueueCheckpointDataNV-pCheckpointDataCount-parameter"></a>
  VUID-vkGetQueueCheckpointDataNV-pCheckpointDataCount-parameter  
  `pCheckpointDataCount` **must** be a valid pointer to a `uint32_t`
  value

- <a href="#VUID-vkGetQueueCheckpointDataNV-pCheckpointData-parameter"
  id="VUID-vkGetQueueCheckpointDataNV-pCheckpointData-parameter"></a>
  VUID-vkGetQueueCheckpointDataNV-pCheckpointData-parameter  
  If the value referenced by `pCheckpointDataCount` is not `0`, and
  `pCheckpointData` is not `NULL`, `pCheckpointData` **must** be a valid
  pointer to an array of `pCheckpointDataCount`
  [VkCheckpointDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCheckpointDataNV.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_diagnostic_checkpoints](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_diagnostic_checkpoints.html),
[VkCheckpointDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCheckpointDataNV.html), [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetQueueCheckpointDataNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
