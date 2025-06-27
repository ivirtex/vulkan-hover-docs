# vkCopyMicromapEXT(3) Manual Page

## Name

vkCopyMicromapEXT - Copy a micromap on the host



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy or compact a micromap on the host, call:

``` c
// Provided by VK_EXT_opacity_micromap
VkResult vkCopyMicromapEXT(
    VkDevice                                    device,
    VkDeferredOperationKHR                      deferredOperation,
    const VkCopyMicromapInfoEXT*                pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns the micromaps.

- `deferredOperation` is an optional
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#deferred-host-operations-requesting"
  target="_blank" rel="noopener">request deferral</a> for this command.

- `pInfo` is a pointer to a
  [VkCopyMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapInfoEXT.html) structure defining
  the copy operation.

## <a href="#_description" class="anchor"></a>Description

This command fulfills the same task as
[vkCmdCopyMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMicromapEXT.html) but is executed by the
host.

Valid Usage

- <a href="#VUID-vkCopyMicromapEXT-deferredOperation-03678"
  id="VUID-vkCopyMicromapEXT-deferredOperation-03678"></a>
  VUID-vkCopyMicromapEXT-deferredOperation-03678  
  Any previous deferred operation that was associated with
  `deferredOperation` **must** be complete

- <a href="#VUID-vkCopyMicromapEXT-buffer-07558"
  id="VUID-vkCopyMicromapEXT-buffer-07558"></a>
  VUID-vkCopyMicromapEXT-buffer-07558  
  The `buffer` used to create `pInfo->src` **must** be bound to
  host-visible device memory

- <a href="#VUID-vkCopyMicromapEXT-buffer-07559"
  id="VUID-vkCopyMicromapEXT-buffer-07559"></a>
  VUID-vkCopyMicromapEXT-buffer-07559  
  The `buffer` used to create `pInfo->dst` **must** be bound to
  host-visible device memory

- <a href="#VUID-vkCopyMicromapEXT-micromapHostCommands-07560"
  id="VUID-vkCopyMicromapEXT-micromapHostCommands-07560"></a>
  VUID-vkCopyMicromapEXT-micromapHostCommands-07560  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-micromapHostCommands"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceOpacityMicromapFeaturesEXT</code>::<code>micromapHostCommands</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCopyMicromapEXT-buffer-07561"
  id="VUID-vkCopyMicromapEXT-buffer-07561"></a>
  VUID-vkCopyMicromapEXT-buffer-07561  
  The `buffer` used to create `pInfo->src` **must** be bound to memory
  that was not allocated with multiple instances

- <a href="#VUID-vkCopyMicromapEXT-buffer-07562"
  id="VUID-vkCopyMicromapEXT-buffer-07562"></a>
  VUID-vkCopyMicromapEXT-buffer-07562  
  The `buffer` used to create `pInfo->dst` **must** be bound to memory
  that was not allocated with multiple instances

Valid Usage (Implicit)

- <a href="#VUID-vkCopyMicromapEXT-device-parameter"
  id="VUID-vkCopyMicromapEXT-device-parameter"></a>
  VUID-vkCopyMicromapEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCopyMicromapEXT-deferredOperation-parameter"
  id="VUID-vkCopyMicromapEXT-deferredOperation-parameter"></a>
  VUID-vkCopyMicromapEXT-deferredOperation-parameter  
  If `deferredOperation` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `deferredOperation` **must** be a valid
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) handle

- <a href="#VUID-vkCopyMicromapEXT-pInfo-parameter"
  id="VUID-vkCopyMicromapEXT-pInfo-parameter"></a>
  VUID-vkCopyMicromapEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkCopyMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapInfoEXT.html) structure

- <a href="#VUID-vkCopyMicromapEXT-deferredOperation-parent"
  id="VUID-vkCopyMicromapEXT-deferredOperation-parent"></a>
  VUID-vkCopyMicromapEXT-deferredOperation-parent  
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

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkCopyMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapInfoEXT.html),
[VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCopyMicromapEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
