# vkCopyAccelerationStructureToMemoryKHR(3) Manual Page

## Name

vkCopyAccelerationStructureToMemoryKHR - Serialize an acceleration
structure on the host



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy an acceleration structure to host accessible memory, call:

``` c
// Provided by VK_KHR_acceleration_structure
VkResult vkCopyAccelerationStructureToMemoryKHR(
    VkDevice                                    device,
    VkDeferredOperationKHR                      deferredOperation,
    const VkCopyAccelerationStructureToMemoryInfoKHR* pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns `pInfo->src`.

- `deferredOperation` is an optional
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#deferred-host-operations-requesting"
  target="_blank" rel="noopener">request deferral</a> for this command.

- `pInfo` is a pointer to a
  [VkCopyAccelerationStructureToMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureToMemoryInfoKHR.html)
  structure defining the copy operation.

## <a href="#_description" class="anchor"></a>Description

This command fulfills the same task as
[vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html)
but is executed by the host.

This command produces the same results as
[vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html),
but writes its result directly to a host pointer, and is executed on the
host rather than the device. The output **may** not necessarily be
bit-for-bit identical, but it can be equally used by either
[vkCmdCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMemoryToAccelerationStructureKHR.html)
or
[vkCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMemoryToAccelerationStructureKHR.html).

Valid Usage

- <a
  href="#VUID-vkCopyAccelerationStructureToMemoryKHR-accelerationStructureHostCommands-03584"
  id="VUID-vkCopyAccelerationStructureToMemoryKHR-accelerationStructureHostCommands-03584"></a>
  VUID-vkCopyAccelerationStructureToMemoryKHR-accelerationStructureHostCommands-03584  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-accelerationStructureHostCommands"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceAccelerationStructureFeaturesKHR</code>::<code>accelerationStructureHostCommands</code></a>
  feature **must** be enabled

<!-- -->

- <a
  href="#VUID-vkCopyAccelerationStructureToMemoryKHR-deferredOperation-03678"
  id="VUID-vkCopyAccelerationStructureToMemoryKHR-deferredOperation-03678"></a>
  VUID-vkCopyAccelerationStructureToMemoryKHR-deferredOperation-03678  
  Any previous deferred operation that was associated with
  `deferredOperation` **must** be complete

- <a href="#VUID-vkCopyAccelerationStructureToMemoryKHR-buffer-03731"
  id="VUID-vkCopyAccelerationStructureToMemoryKHR-buffer-03731"></a>
  VUID-vkCopyAccelerationStructureToMemoryKHR-buffer-03731  
  The `buffer` used to create `pInfo->src` **must** be bound to
  host-visible device memory

- <a href="#VUID-vkCopyAccelerationStructureToMemoryKHR-pInfo-03732"
  id="VUID-vkCopyAccelerationStructureToMemoryKHR-pInfo-03732"></a>
  VUID-vkCopyAccelerationStructureToMemoryKHR-pInfo-03732  
  `pInfo->dst.hostAddress` **must** be a valid host pointer

- <a href="#VUID-vkCopyAccelerationStructureToMemoryKHR-pInfo-03751"
  id="VUID-vkCopyAccelerationStructureToMemoryKHR-pInfo-03751"></a>
  VUID-vkCopyAccelerationStructureToMemoryKHR-pInfo-03751  
  `pInfo->dst.hostAddress` **must** be aligned to 16 bytes

- <a href="#VUID-vkCopyAccelerationStructureToMemoryKHR-buffer-03783"
  id="VUID-vkCopyAccelerationStructureToMemoryKHR-buffer-03783"></a>
  VUID-vkCopyAccelerationStructureToMemoryKHR-buffer-03783  
  The `buffer` used to create `pInfo->src` **must** be bound to memory
  that was not allocated with multiple instances

Valid Usage (Implicit)

- <a href="#VUID-vkCopyAccelerationStructureToMemoryKHR-device-parameter"
  id="VUID-vkCopyAccelerationStructureToMemoryKHR-device-parameter"></a>
  VUID-vkCopyAccelerationStructureToMemoryKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkCopyAccelerationStructureToMemoryKHR-deferredOperation-parameter"
  id="VUID-vkCopyAccelerationStructureToMemoryKHR-deferredOperation-parameter"></a>
  VUID-vkCopyAccelerationStructureToMemoryKHR-deferredOperation-parameter  
  If `deferredOperation` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `deferredOperation` **must** be a valid
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) handle

- <a href="#VUID-vkCopyAccelerationStructureToMemoryKHR-pInfo-parameter"
  id="VUID-vkCopyAccelerationStructureToMemoryKHR-pInfo-parameter"></a>
  VUID-vkCopyAccelerationStructureToMemoryKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkCopyAccelerationStructureToMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureToMemoryInfoKHR.html)
  structure

- <a
  href="#VUID-vkCopyAccelerationStructureToMemoryKHR-deferredOperation-parent"
  id="VUID-vkCopyAccelerationStructureToMemoryKHR-deferredOperation-parent"></a>
  VUID-vkCopyAccelerationStructureToMemoryKHR-deferredOperation-parent  
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
[VkCopyAccelerationStructureToMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureToMemoryInfoKHR.html),
[VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCopyAccelerationStructureToMemoryKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
