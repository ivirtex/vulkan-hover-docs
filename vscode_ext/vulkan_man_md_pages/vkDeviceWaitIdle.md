# vkDeviceWaitIdle(3) Manual Page

## Name

vkDeviceWaitIdle - Wait for a device to become idle



## <a href="#_c_specification" class="anchor"></a>C Specification

To wait on the host for the completion of outstanding queue operations
for all queues on a given logical device, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkDeviceWaitIdle(
    VkDevice                                    device);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device to idle.

## <a href="#_description" class="anchor"></a>Description

`vkDeviceWaitIdle` is equivalent to calling `vkQueueWaitIdle` for all
queues owned by `device`.

Valid Usage (Implicit)

- <a href="#VUID-vkDeviceWaitIdle-device-parameter"
  id="VUID-vkDeviceWaitIdle-device-parameter"></a>
  VUID-vkDeviceWaitIdle-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

Host Synchronization

- Host access to all `VkQueue` objects created from `device` **must** be
  externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_DEVICE_LOST`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDeviceWaitIdle"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
