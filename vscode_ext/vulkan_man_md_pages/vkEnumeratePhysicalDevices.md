# vkEnumeratePhysicalDevices(3) Manual Page

## Name

vkEnumeratePhysicalDevices - Enumerates the physical devices accessible
to a Vulkan instance



## <a href="#_c_specification" class="anchor"></a>C Specification

To retrieve a list of physical device objects representing the physical
devices installed in the system, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkEnumeratePhysicalDevices(
    VkInstance                                  instance,
    uint32_t*                                   pPhysicalDeviceCount,
    VkPhysicalDevice*                           pPhysicalDevices);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is a handle to a Vulkan instance previously created with
  [vkCreateInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateInstance.html).

- `pPhysicalDeviceCount` is a pointer to an integer related to the
  number of physical devices available or queried, as described below.

- `pPhysicalDevices` is either `NULL` or a pointer to an array of
  `VkPhysicalDevice` handles.

## <a href="#_description" class="anchor"></a>Description

If `pPhysicalDevices` is `NULL`, then the number of physical devices
available is returned in `pPhysicalDeviceCount`. Otherwise,
`pPhysicalDeviceCount` **must** point to a variable set by the
application to the number of elements in the `pPhysicalDevices` array,
and on return the variable is overwritten with the number of handles
actually written to `pPhysicalDevices`. If `pPhysicalDeviceCount` is
less than the number of physical devices available, at most
`pPhysicalDeviceCount` structures will be written, and `VK_INCOMPLETE`
will be returned instead of `VK_SUCCESS`, to indicate that not all the
available physical devices were returned.

Valid Usage (Implicit)

- <a href="#VUID-vkEnumeratePhysicalDevices-instance-parameter"
  id="VUID-vkEnumeratePhysicalDevices-instance-parameter"></a>
  VUID-vkEnumeratePhysicalDevices-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a
  href="#VUID-vkEnumeratePhysicalDevices-pPhysicalDeviceCount-parameter"
  id="VUID-vkEnumeratePhysicalDevices-pPhysicalDeviceCount-parameter"></a>
  VUID-vkEnumeratePhysicalDevices-pPhysicalDeviceCount-parameter  
  `pPhysicalDeviceCount` **must** be a valid pointer to a `uint32_t`
  value

- <a href="#VUID-vkEnumeratePhysicalDevices-pPhysicalDevices-parameter"
  id="VUID-vkEnumeratePhysicalDevices-pPhysicalDevices-parameter"></a>
  VUID-vkEnumeratePhysicalDevices-pPhysicalDevices-parameter  
  If the value referenced by `pPhysicalDeviceCount` is not `0`, and
  `pPhysicalDevices` is not `NULL`, `pPhysicalDevices` **must** be a
  valid pointer to an array of `pPhysicalDeviceCount`
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handles

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INITIALIZATION_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkEnumeratePhysicalDevices"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
