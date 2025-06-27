# vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR(3) Manual Page

## Name

vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR - Returns properties
describing what cooperative matrix types are supported



## <a href="#_c_specification" class="anchor"></a>C Specification

To enumerate the supported cooperative matrix types and operations,
call:

``` c
// Provided by VK_KHR_cooperative_matrix
VkResult vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pPropertyCount,
    VkCooperativeMatrixPropertiesKHR*           pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device.

- `pPropertyCount` is a pointer to an integer related to the number of
  cooperative matrix properties available or queried.

- `pProperties` is either `NULL` or a pointer to an array of
  [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)
  structures.

## <a href="#_description" class="anchor"></a>Description

If `pProperties` is `NULL`, then the number of cooperative matrix
properties available is returned in `pPropertyCount`. Otherwise,
`pPropertyCount` **must** point to a variable set by the application to
the number of elements in the `pProperties` array, and on return the
variable is overwritten with the number of structures actually written
to `pProperties`. If `pPropertyCount` is less than the number of
cooperative matrix properties available, at most `pPropertyCount`
structures will be written, and `VK_INCOMPLETE` will be returned instead
of `VK_SUCCESS`, to indicate that not all the available cooperative
matrix properties were returned.

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR-pPropertyCount-parameter"
  id="VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR-pPropertyCount-parameter"></a>
  VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value

- <a
  href="#VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR-pProperties-parameter"
  id="VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR-pProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and
  `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer
  to an array of `pPropertyCount`
  [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)
  structures

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_cooperative_matrix](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_cooperative_matrix.html),
[VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
