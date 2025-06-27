# vkEnumeratePhysicalDeviceGroups(3) Manual Page

## Name

vkEnumeratePhysicalDeviceGroups - Enumerates groups of physical devices
that can be used to create a single logical device



## <a href="#_c_specification" class="anchor"></a>C Specification

To retrieve a list of the device groups present in the system, call:

``` c
// Provided by VK_VERSION_1_1
VkResult vkEnumeratePhysicalDeviceGroups(
    VkInstance                                  instance,
    uint32_t*                                   pPhysicalDeviceGroupCount,
    VkPhysicalDeviceGroupProperties*            pPhysicalDeviceGroupProperties);
```

or the equivalent command

``` c
// Provided by VK_KHR_device_group_creation
VkResult vkEnumeratePhysicalDeviceGroupsKHR(
    VkInstance                                  instance,
    uint32_t*                                   pPhysicalDeviceGroupCount,
    VkPhysicalDeviceGroupProperties*            pPhysicalDeviceGroupProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is a handle to a Vulkan instance previously created with
  [vkCreateInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateInstance.html).

- `pPhysicalDeviceGroupCount` is a pointer to an integer related to the
  number of device groups available or queried, as described below.

- `pPhysicalDeviceGroupProperties` is either `NULL` or a pointer to an
  array of
  [VkPhysicalDeviceGroupProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceGroupProperties.html)
  structures.

## <a href="#_description" class="anchor"></a>Description

If `pPhysicalDeviceGroupProperties` is `NULL`, then the number of device
groups available is returned in `pPhysicalDeviceGroupCount`. Otherwise,
`pPhysicalDeviceGroupCount` **must** point to a variable set by the
application to the number of elements in the
`pPhysicalDeviceGroupProperties` array, and on return the variable is
overwritten with the number of structures actually written to
`pPhysicalDeviceGroupProperties`. If `pPhysicalDeviceGroupCount` is less
than the number of device groups available, at most
`pPhysicalDeviceGroupCount` structures will be written, and
`VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate
that not all the available device groups were returned.

Every physical device **must** be in exactly one device group.

Valid Usage (Implicit)

- <a href="#VUID-vkEnumeratePhysicalDeviceGroups-instance-parameter"
  id="VUID-vkEnumeratePhysicalDeviceGroups-instance-parameter"></a>
  VUID-vkEnumeratePhysicalDeviceGroups-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a
  href="#VUID-vkEnumeratePhysicalDeviceGroups-pPhysicalDeviceGroupCount-parameter"
  id="VUID-vkEnumeratePhysicalDeviceGroups-pPhysicalDeviceGroupCount-parameter"></a>
  VUID-vkEnumeratePhysicalDeviceGroups-pPhysicalDeviceGroupCount-parameter  
  `pPhysicalDeviceGroupCount` **must** be a valid pointer to a
  `uint32_t` value

- <a
  href="#VUID-vkEnumeratePhysicalDeviceGroups-pPhysicalDeviceGroupProperties-parameter"
  id="VUID-vkEnumeratePhysicalDeviceGroups-pPhysicalDeviceGroupProperties-parameter"></a>
  VUID-vkEnumeratePhysicalDeviceGroups-pPhysicalDeviceGroupProperties-parameter  
  If the value referenced by `pPhysicalDeviceGroupCount` is not `0`, and
  `pPhysicalDeviceGroupProperties` is not `NULL`,
  `pPhysicalDeviceGroupProperties` **must** be a valid pointer to an
  array of `pPhysicalDeviceGroupCount`
  [VkPhysicalDeviceGroupProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceGroupProperties.html)
  structures

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INITIALIZATION_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html),
[VkPhysicalDeviceGroupProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceGroupProperties.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkEnumeratePhysicalDeviceGroups"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
