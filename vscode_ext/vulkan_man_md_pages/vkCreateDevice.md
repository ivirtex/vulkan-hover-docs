# vkCreateDevice(3) Manual Page

## Name

vkCreateDevice - Create a new device instance



## <a href="#_c_specification" class="anchor"></a>C Specification

A logical device is created as a *connection* to a physical device. To
create a logical device, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateDevice(
    VkPhysicalDevice                            physicalDevice,
    const VkDeviceCreateInfo*                   pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDevice*                                   pDevice);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` **must** be one of the device handles returned from a
  call to `vkEnumeratePhysicalDevices` (see <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-physical-device-enumeration"
  target="_blank" rel="noopener">Physical Device Enumeration</a>).

- `pCreateInfo` is a pointer to a
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) structure containing
  information about how to create the device.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pDevice` is a pointer to a handle in which the created
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) is returned.

## <a href="#_description" class="anchor"></a>Description

`vkCreateDevice` verifies that extensions and features requested in the
`ppEnabledExtensionNames` and `pEnabledFeatures` members of
`pCreateInfo`, respectively, are supported by the implementation. If any
requested extension is not supported, `vkCreateDevice` **must** return
`VK_ERROR_EXTENSION_NOT_PRESENT`. If any requested feature is not
supported, `vkCreateDevice` **must** return
`VK_ERROR_FEATURE_NOT_PRESENT`. Support for extensions **can** be
checked before creating a device by querying
[vkEnumerateDeviceExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateDeviceExtensionProperties.html).
Support for features **can** similarly be checked by querying
[vkGetPhysicalDeviceFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures.html).

After verifying and enabling the extensions the `VkDevice` object is
created and returned to the application.

Multiple logical devices **can** be created from the same physical
device. Logical device creation **may** fail due to lack of
device-specific resources (in addition to other errors). If that occurs,
`vkCreateDevice` will return `VK_ERROR_TOO_MANY_OBJECTS`.

Valid Usage

- <a href="#VUID-vkCreateDevice-ppEnabledExtensionNames-01387"
  id="VUID-vkCreateDevice-ppEnabledExtensionNames-01387"></a>
  VUID-vkCreateDevice-ppEnabledExtensionNames-01387  
  All <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-extensions-extensiondependencies"
  target="_blank" rel="noopener">required device extensions</a> for each
  extension in the
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html)::`ppEnabledExtensionNames`
  list **must** also be present in that list

Valid Usage (Implicit)

- <a href="#VUID-vkCreateDevice-physicalDevice-parameter"
  id="VUID-vkCreateDevice-physicalDevice-parameter"></a>
  VUID-vkCreateDevice-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkCreateDevice-pCreateInfo-parameter"
  id="VUID-vkCreateDevice-pCreateInfo-parameter"></a>
  VUID-vkCreateDevice-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) structure

- <a href="#VUID-vkCreateDevice-pAllocator-parameter"
  id="VUID-vkCreateDevice-pAllocator-parameter"></a>
  VUID-vkCreateDevice-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateDevice-pDevice-parameter"
  id="VUID-vkCreateDevice-pDevice-parameter"></a>
  VUID-vkCreateDevice-pDevice-parameter  
  `pDevice` **must** be a valid pointer to a [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)
  handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INITIALIZATION_FAILED`

- `VK_ERROR_EXTENSION_NOT_PRESENT`

- `VK_ERROR_FEATURE_NOT_PRESENT`

- `VK_ERROR_TOO_MANY_OBJECTS`

- `VK_ERROR_DEVICE_LOST`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateDevice"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
