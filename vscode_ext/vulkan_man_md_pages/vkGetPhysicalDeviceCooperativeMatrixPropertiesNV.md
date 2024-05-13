# vkGetPhysicalDeviceCooperativeMatrixPropertiesNV(3) Manual Page

## Name

vkGetPhysicalDeviceCooperativeMatrixPropertiesNV - Returns properties
describing what cooperative matrix types are supported



## <a href="#_c_specification" class="anchor"></a>C Specification

To enumerate the supported cooperative matrix types and operations,
call:

``` c
// Provided by VK_NV_cooperative_matrix
VkResult vkGetPhysicalDeviceCooperativeMatrixPropertiesNV(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pPropertyCount,
    VkCooperativeMatrixPropertiesNV*            pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device.

- `pPropertyCount` is a pointer to an integer related to the number of
  cooperative matrix properties available or queried.

- `pProperties` is either `NULL` or a pointer to an array of
  [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html)
  structures.

## <a href="#_description" class="anchor"></a>Description

If `pProperties` is `NULL`, then the number of cooperative matrix
properties available is returned in `pPropertyCount`. Otherwise,
`pPropertyCount` **must** point to a variable set by the user to the
number of elements in the `pProperties` array, and on return the
variable is overwritten with the number of structures actually written
to `pProperties`. If `pPropertyCount` is less than the number of
cooperative matrix properties available, at most `pPropertyCount`
structures will be written, and `VK_INCOMPLETE` will be returned instead
of `VK_SUCCESS`, to indicate that not all the available cooperative
matrix properties were returned.

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesNV-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesNV-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesNV-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesNV-pPropertyCount-parameter"
  id="VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesNV-pPropertyCount-parameter"></a>
  VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesNV-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value

- <a
  href="#VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesNV-pProperties-parameter"
  id="VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesNV-pProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceCooperativeMatrixPropertiesNV-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and
  `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer
  to an array of `pPropertyCount`
  [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html)
  structures

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_cooperative_matrix](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_cooperative_matrix.html),
[VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceCooperativeMatrixPropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
