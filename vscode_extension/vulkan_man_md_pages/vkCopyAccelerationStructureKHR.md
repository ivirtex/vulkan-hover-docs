# vkCopyAccelerationStructureKHR(3) Manual Page

## Name

vkCopyAccelerationStructureKHR - Copy an acceleration structure on the
host



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy or compact an acceleration structure on the host, call:

``` c
// Provided by VK_KHR_acceleration_structure
VkResult vkCopyAccelerationStructureKHR(
    VkDevice                                    device,
    VkDeferredOperationKHR                      deferredOperation,
    const VkCopyAccelerationStructureInfoKHR*   pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns the acceleration structures.

- `deferredOperation` is an optional
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#deferred-host-operations-requesting"
  target="_blank" rel="noopener">request deferral</a> for this command.

- `pInfo` is a pointer to a
  [VkCopyAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureInfoKHR.html)
  structure defining the copy operation.

## <a href="#_description" class="anchor"></a>Description

This command fulfills the same task as
[vkCmdCopyAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureKHR.html)
but is executed by the host.

Valid Usage

- <a
  href="#VUID-vkCopyAccelerationStructureKHR-accelerationStructureHostCommands-03582"
  id="VUID-vkCopyAccelerationStructureKHR-accelerationStructureHostCommands-03582"></a>
  VUID-vkCopyAccelerationStructureKHR-accelerationStructureHostCommands-03582  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-accelerationStructureHostCommands"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceAccelerationStructureFeaturesKHR</code>::<code>accelerationStructureHostCommands</code></a>
  feature **must** be enabled

<!-- -->

- <a href="#VUID-vkCopyAccelerationStructureKHR-deferredOperation-03678"
  id="VUID-vkCopyAccelerationStructureKHR-deferredOperation-03678"></a>
  VUID-vkCopyAccelerationStructureKHR-deferredOperation-03678  
  Any previous deferred operation that was associated with
  `deferredOperation` **must** be complete

- <a href="#VUID-vkCopyAccelerationStructureKHR-buffer-03727"
  id="VUID-vkCopyAccelerationStructureKHR-buffer-03727"></a>
  VUID-vkCopyAccelerationStructureKHR-buffer-03727  
  The `buffer` used to create `pInfo->src` **must** be bound to
  host-visible device memory

- <a href="#VUID-vkCopyAccelerationStructureKHR-buffer-03728"
  id="VUID-vkCopyAccelerationStructureKHR-buffer-03728"></a>
  VUID-vkCopyAccelerationStructureKHR-buffer-03728  
  The `buffer` used to create `pInfo->dst` **must** be bound to
  host-visible device memory

- <a href="#VUID-vkCopyAccelerationStructureKHR-buffer-03780"
  id="VUID-vkCopyAccelerationStructureKHR-buffer-03780"></a>
  VUID-vkCopyAccelerationStructureKHR-buffer-03780  
  The `buffer` used to create `pInfo->src` **must** be bound to memory
  that was not allocated with multiple instances

- <a href="#VUID-vkCopyAccelerationStructureKHR-buffer-03781"
  id="VUID-vkCopyAccelerationStructureKHR-buffer-03781"></a>
  VUID-vkCopyAccelerationStructureKHR-buffer-03781  
  The `buffer` used to create `pInfo->dst` **must** be bound to memory
  that was not allocated with multiple instances

Valid Usage (Implicit)

- <a href="#VUID-vkCopyAccelerationStructureKHR-device-parameter"
  id="VUID-vkCopyAccelerationStructureKHR-device-parameter"></a>
  VUID-vkCopyAccelerationStructureKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkCopyAccelerationStructureKHR-deferredOperation-parameter"
  id="VUID-vkCopyAccelerationStructureKHR-deferredOperation-parameter"></a>
  VUID-vkCopyAccelerationStructureKHR-deferredOperation-parameter  
  If `deferredOperation` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `deferredOperation` **must** be a valid
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) handle

- <a href="#VUID-vkCopyAccelerationStructureKHR-pInfo-parameter"
  id="VUID-vkCopyAccelerationStructureKHR-pInfo-parameter"></a>
  VUID-vkCopyAccelerationStructureKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkCopyAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureInfoKHR.html)
  structure

- <a href="#VUID-vkCopyAccelerationStructureKHR-deferredOperation-parent"
  id="VUID-vkCopyAccelerationStructureKHR-deferredOperation-parent"></a>
  VUID-vkCopyAccelerationStructureKHR-deferredOperation-parent  
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
[VkCopyAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureInfoKHR.html),
[VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCopyAccelerationStructureKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
