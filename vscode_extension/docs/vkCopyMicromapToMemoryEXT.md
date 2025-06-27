# vkCopyMicromapToMemoryEXT(3) Manual Page

## Name

vkCopyMicromapToMemoryEXT - Serialize a micromap on the host



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy a micromap to host accessible memory, call:

``` c
// Provided by VK_EXT_opacity_micromap
VkResult vkCopyMicromapToMemoryEXT(
    VkDevice                                    device,
    VkDeferredOperationKHR                      deferredOperation,
    const VkCopyMicromapToMemoryInfoEXT*        pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns `pInfo->src`.

- `deferredOperation` is an optional
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#deferred-host-operations-requesting"
  target="_blank" rel="noopener">request deferral</a> for this command.

- `pInfo` is a pointer to a
  [VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapToMemoryInfoEXT.html)
  structure defining the copy operation.

## <a href="#_description" class="anchor"></a>Description

This command fulfills the same task as
[vkCmdCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMicromapToMemoryEXT.html) but is
executed by the host.

This command produces the same results as
[vkCmdCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMicromapToMemoryEXT.html), but
writes its result directly to a host pointer, and is executed on the
host rather than the device. The output **may** not necessarily be
bit-for-bit identical, but it can be equally used by either
[vkCmdCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMemoryToMicromapEXT.html) or
[vkCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMemoryToMicromapEXT.html).

Valid Usage

- <a href="#VUID-vkCopyMicromapToMemoryEXT-deferredOperation-03678"
  id="VUID-vkCopyMicromapToMemoryEXT-deferredOperation-03678"></a>
  VUID-vkCopyMicromapToMemoryEXT-deferredOperation-03678  
  Any previous deferred operation that was associated with
  `deferredOperation` **must** be complete

- <a href="#VUID-vkCopyMicromapToMemoryEXT-buffer-07568"
  id="VUID-vkCopyMicromapToMemoryEXT-buffer-07568"></a>
  VUID-vkCopyMicromapToMemoryEXT-buffer-07568  
  The `buffer` used to create `pInfo->src` **must** be bound to
  host-visible device memory

- <a href="#VUID-vkCopyMicromapToMemoryEXT-pInfo-07569"
  id="VUID-vkCopyMicromapToMemoryEXT-pInfo-07569"></a>
  VUID-vkCopyMicromapToMemoryEXT-pInfo-07569  
  `pInfo->dst.hostAddress` **must** be a valid host pointer

- <a href="#VUID-vkCopyMicromapToMemoryEXT-pInfo-07570"
  id="VUID-vkCopyMicromapToMemoryEXT-pInfo-07570"></a>
  VUID-vkCopyMicromapToMemoryEXT-pInfo-07570  
  `pInfo->dst.hostAddress` **must** be aligned to 16 bytes

- <a href="#VUID-vkCopyMicromapToMemoryEXT-micromapHostCommands-07571"
  id="VUID-vkCopyMicromapToMemoryEXT-micromapHostCommands-07571"></a>
  VUID-vkCopyMicromapToMemoryEXT-micromapHostCommands-07571  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-micromapHostCommands"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceOpacityMicromapFeaturesEXT</code>::<code>micromapHostCommands</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCopyMicromapToMemoryEXT-buffer-07572"
  id="VUID-vkCopyMicromapToMemoryEXT-buffer-07572"></a>
  VUID-vkCopyMicromapToMemoryEXT-buffer-07572  
  The `buffer` used to create `pInfo->src` **must** be bound to memory
  that was not allocated with multiple instances

Valid Usage (Implicit)

- <a href="#VUID-vkCopyMicromapToMemoryEXT-device-parameter"
  id="VUID-vkCopyMicromapToMemoryEXT-device-parameter"></a>
  VUID-vkCopyMicromapToMemoryEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCopyMicromapToMemoryEXT-deferredOperation-parameter"
  id="VUID-vkCopyMicromapToMemoryEXT-deferredOperation-parameter"></a>
  VUID-vkCopyMicromapToMemoryEXT-deferredOperation-parameter  
  If `deferredOperation` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `deferredOperation` **must** be a valid
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) handle

- <a href="#VUID-vkCopyMicromapToMemoryEXT-pInfo-parameter"
  id="VUID-vkCopyMicromapToMemoryEXT-pInfo-parameter"></a>
  VUID-vkCopyMicromapToMemoryEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapToMemoryInfoEXT.html)
  structure

- <a href="#VUID-vkCopyMicromapToMemoryEXT-deferredOperation-parent"
  id="VUID-vkCopyMicromapToMemoryEXT-deferredOperation-parent"></a>
  VUID-vkCopyMicromapToMemoryEXT-deferredOperation-parent  
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
[VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapToMemoryInfoEXT.html),
[VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCopyMicromapToMemoryEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
