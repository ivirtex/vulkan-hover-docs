# vkGetFenceStatus(3) Manual Page

## Name

vkGetFenceStatus - Return the status of a fence



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the status of a fence from the host, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkGetFenceStatus(
    VkDevice                                    device,
    VkFence                                     fence);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the fence.

- `fence` is the handle of the fence to query.

## <a href="#_description" class="anchor"></a>Description

Upon success, `vkGetFenceStatus` returns the status of the fence object,
with the following return codes:

| Status                 | Meaning                                                                                                    |
|------------------------|------------------------------------------------------------------------------------------------------------|
| `VK_SUCCESS`           | The fence specified by `fence` is signaled.                                                                |
| `VK_NOT_READY`         | The fence specified by `fence` is unsignaled.                                                              |
| `VK_ERROR_DEVICE_LOST` | The device has been lost. See <a                                                                           
                          href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-lost-device"  
                          target="_blank" rel="noopener">Lost Device</a>.                                                             |

Table 1. Fence Object Status Codes

If a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-submission"
target="_blank" rel="noopener">queue submission</a> command is pending
execution, then the value returned by this command **may** immediately
be out of date.

If the device has been lost (see <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-lost-device"
target="_blank" rel="noopener">Lost Device</a>), `vkGetFenceStatus`
**may** return any of the above status codes. If the device has been
lost and `vkGetFenceStatus` is called repeatedly, it will eventually
return either `VK_SUCCESS` or `VK_ERROR_DEVICE_LOST`.

Valid Usage (Implicit)

- <a href="#VUID-vkGetFenceStatus-device-parameter"
  id="VUID-vkGetFenceStatus-device-parameter"></a>
  VUID-vkGetFenceStatus-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetFenceStatus-fence-parameter"
  id="VUID-vkGetFenceStatus-fence-parameter"></a>
  VUID-vkGetFenceStatus-fence-parameter  
  `fence` **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html) handle

- <a href="#VUID-vkGetFenceStatus-fence-parent"
  id="VUID-vkGetFenceStatus-fence-parent"></a>
  VUID-vkGetFenceStatus-fence-parent  
  `fence` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_NOT_READY`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_DEVICE_LOST`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetFenceStatus"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
