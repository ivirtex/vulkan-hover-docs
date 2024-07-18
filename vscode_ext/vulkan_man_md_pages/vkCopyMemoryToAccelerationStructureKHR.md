# vkCopyMemoryToAccelerationStructureKHR(3) Manual Page

## Name

vkCopyMemoryToAccelerationStructureKHR - Deserialize an acceleration
structure on the host



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy host accessible memory to an acceleration structure, call:

``` c
// Provided by VK_KHR_acceleration_structure
VkResult vkCopyMemoryToAccelerationStructureKHR(
    VkDevice                                    device,
    VkDeferredOperationKHR                      deferredOperation,
    const VkCopyMemoryToAccelerationStructureInfoKHR* pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns `pInfo->dst`.

- `deferredOperation` is an optional
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#deferred-host-operations-requesting"
  target="_blank" rel="noopener">request deferral</a> for this command.

- `pInfo` is a pointer to a
  [VkCopyMemoryToAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToAccelerationStructureInfoKHR.html)
  structure defining the copy operation.

## <a href="#_description" class="anchor"></a>Description

This command fulfills the same task as
[vkCmdCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMemoryToAccelerationStructureKHR.html)
but is executed by the host.

This command can accept acceleration structures produced by either
[vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html)
or
[vkCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyAccelerationStructureToMemoryKHR.html).

Valid Usage

- <a
  href="#VUID-vkCopyMemoryToAccelerationStructureKHR-accelerationStructureHostCommands-03583"
  id="VUID-vkCopyMemoryToAccelerationStructureKHR-accelerationStructureHostCommands-03583"></a>
  VUID-vkCopyMemoryToAccelerationStructureKHR-accelerationStructureHostCommands-03583  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-accelerationStructureHostCommands"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceAccelerationStructureFeaturesKHR</code>::<code>accelerationStructureHostCommands</code></a>
  feature **must** be enabled

<!-- -->

- <a
  href="#VUID-vkCopyMemoryToAccelerationStructureKHR-deferredOperation-03678"
  id="VUID-vkCopyMemoryToAccelerationStructureKHR-deferredOperation-03678"></a>
  VUID-vkCopyMemoryToAccelerationStructureKHR-deferredOperation-03678  
  Any previous deferred operation that was associated with
  `deferredOperation` **must** be complete

- <a href="#VUID-vkCopyMemoryToAccelerationStructureKHR-pInfo-03729"
  id="VUID-vkCopyMemoryToAccelerationStructureKHR-pInfo-03729"></a>
  VUID-vkCopyMemoryToAccelerationStructureKHR-pInfo-03729  
  `pInfo->src.hostAddress` **must** be a valid host pointer

- <a href="#VUID-vkCopyMemoryToAccelerationStructureKHR-pInfo-03750"
  id="VUID-vkCopyMemoryToAccelerationStructureKHR-pInfo-03750"></a>
  VUID-vkCopyMemoryToAccelerationStructureKHR-pInfo-03750  
  `pInfo->src.hostAddress` **must** be aligned to 16 bytes

- <a href="#VUID-vkCopyMemoryToAccelerationStructureKHR-buffer-03730"
  id="VUID-vkCopyMemoryToAccelerationStructureKHR-buffer-03730"></a>
  VUID-vkCopyMemoryToAccelerationStructureKHR-buffer-03730  
  The `buffer` used to create `pInfo->dst` **must** be bound to
  host-visible device memory

- <a href="#VUID-vkCopyMemoryToAccelerationStructureKHR-buffer-03782"
  id="VUID-vkCopyMemoryToAccelerationStructureKHR-buffer-03782"></a>
  VUID-vkCopyMemoryToAccelerationStructureKHR-buffer-03782  
  The `buffer` used to create `pInfo->dst` **must** be bound to memory
  that was not allocated with multiple instances

Valid Usage (Implicit)

- <a href="#VUID-vkCopyMemoryToAccelerationStructureKHR-device-parameter"
  id="VUID-vkCopyMemoryToAccelerationStructureKHR-device-parameter"></a>
  VUID-vkCopyMemoryToAccelerationStructureKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkCopyMemoryToAccelerationStructureKHR-deferredOperation-parameter"
  id="VUID-vkCopyMemoryToAccelerationStructureKHR-deferredOperation-parameter"></a>
  VUID-vkCopyMemoryToAccelerationStructureKHR-deferredOperation-parameter  
  If `deferredOperation` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `deferredOperation` **must** be a valid
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) handle

- <a href="#VUID-vkCopyMemoryToAccelerationStructureKHR-pInfo-parameter"
  id="VUID-vkCopyMemoryToAccelerationStructureKHR-pInfo-parameter"></a>
  VUID-vkCopyMemoryToAccelerationStructureKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkCopyMemoryToAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToAccelerationStructureInfoKHR.html)
  structure

- <a
  href="#VUID-vkCopyMemoryToAccelerationStructureKHR-deferredOperation-parent"
  id="VUID-vkCopyMemoryToAccelerationStructureKHR-deferredOperation-parent"></a>
  VUID-vkCopyMemoryToAccelerationStructureKHR-deferredOperation-parent  
  If `deferredOperation` is a valid handle, it **must** have been
  created, allocated, or retrieved from `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_OPERATION_DEFERRED_KHR`

- `VK_OPERATION_NOT_DEFERRED_KHR`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkCopyMemoryToAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToAccelerationStructureInfoKHR.html),
[VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCopyMemoryToAccelerationStructureKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
