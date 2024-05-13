# vkCopyMemoryToMicromapEXT(3) Manual Page

## Name

vkCopyMemoryToMicromapEXT - Deserialize a micromap on the host



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy host accessible memory to a micromap, call:

``` c
// Provided by VK_EXT_opacity_micromap
VkResult vkCopyMemoryToMicromapEXT(
    VkDevice                                    device,
    VkDeferredOperationKHR                      deferredOperation,
    const VkCopyMemoryToMicromapInfoEXT*        pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns `pInfo->dst`.

- `deferredOperation` is an optional
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#deferred-host-operations-requesting"
  target="_blank" rel="noopener">request deferral</a> for this command.

- `pInfo` is a pointer to a
  [VkCopyMemoryToMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToMicromapInfoEXT.html)
  structure defining the copy operation.

## <a href="#_description" class="anchor"></a>Description

This command fulfills the same task as
[vkCmdCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMemoryToMicromapEXT.html) but is
executed by the host.

This command can accept micromaps produced by either
[vkCmdCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMicromapToMemoryEXT.html) or
[vkCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMicromapToMemoryEXT.html).

Valid Usage

- <a href="#VUID-vkCopyMemoryToMicromapEXT-deferredOperation-03678"
  id="VUID-vkCopyMemoryToMicromapEXT-deferredOperation-03678"></a>
  VUID-vkCopyMemoryToMicromapEXT-deferredOperation-03678  
  Any previous deferred operation that was associated with
  `deferredOperation` **must** be complete

- <a href="#VUID-vkCopyMemoryToMicromapEXT-pInfo-07563"
  id="VUID-vkCopyMemoryToMicromapEXT-pInfo-07563"></a>
  VUID-vkCopyMemoryToMicromapEXT-pInfo-07563  
  `pInfo->src.hostAddress` **must** be a valid host pointer

- <a href="#VUID-vkCopyMemoryToMicromapEXT-pInfo-07564"
  id="VUID-vkCopyMemoryToMicromapEXT-pInfo-07564"></a>
  VUID-vkCopyMemoryToMicromapEXT-pInfo-07564  
  `pInfo->src.hostAddress` **must** be aligned to 16 bytes

- <a href="#VUID-vkCopyMemoryToMicromapEXT-buffer-07565"
  id="VUID-vkCopyMemoryToMicromapEXT-buffer-07565"></a>
  VUID-vkCopyMemoryToMicromapEXT-buffer-07565  
  The `buffer` used to create `pInfo->dst` **must** be bound to
  host-visible device memory

- <a href="#VUID-vkCopyMemoryToMicromapEXT-micromapHostCommands-07566"
  id="VUID-vkCopyMemoryToMicromapEXT-micromapHostCommands-07566"></a>
  VUID-vkCopyMemoryToMicromapEXT-micromapHostCommands-07566  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-micromapHostCommands"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceOpacityMicromapFeaturesEXT</code>::<code>micromapHostCommands</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCopyMemoryToMicromapEXT-buffer-07567"
  id="VUID-vkCopyMemoryToMicromapEXT-buffer-07567"></a>
  VUID-vkCopyMemoryToMicromapEXT-buffer-07567  
  The `buffer` used to create `pInfo->dst` **must** be bound to memory
  that was not allocated with multiple instances

Valid Usage (Implicit)

- <a href="#VUID-vkCopyMemoryToMicromapEXT-device-parameter"
  id="VUID-vkCopyMemoryToMicromapEXT-device-parameter"></a>
  VUID-vkCopyMemoryToMicromapEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCopyMemoryToMicromapEXT-deferredOperation-parameter"
  id="VUID-vkCopyMemoryToMicromapEXT-deferredOperation-parameter"></a>
  VUID-vkCopyMemoryToMicromapEXT-deferredOperation-parameter  
  If `deferredOperation` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `deferredOperation` **must** be a valid
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) handle

- <a href="#VUID-vkCopyMemoryToMicromapEXT-pInfo-parameter"
  id="VUID-vkCopyMemoryToMicromapEXT-pInfo-parameter"></a>
  VUID-vkCopyMemoryToMicromapEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkCopyMemoryToMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToMicromapInfoEXT.html)
  structure

- <a href="#VUID-vkCopyMemoryToMicromapEXT-deferredOperation-parent"
  id="VUID-vkCopyMemoryToMicromapEXT-deferredOperation-parent"></a>
  VUID-vkCopyMemoryToMicromapEXT-deferredOperation-parent  
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
[VkCopyMemoryToMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToMicromapInfoEXT.html),
[VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCopyMemoryToMicromapEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
