# VkPhysicalDeviceDepthBiasControlFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceDepthBiasControlFeaturesEXT - Structure indicating support for depth bias scaling and representation control



## [](#_c_specification)C Specification

The `VkPhysicalDeviceDepthBiasControlFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_depth_bias_control
typedef struct VkPhysicalDeviceDepthBiasControlFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           depthBiasControl;
    VkBool32           leastRepresentableValueForceUnormRepresentation;
    VkBool32           floatRepresentation;
    VkBool32           depthBiasExact;
} VkPhysicalDeviceDepthBiasControlFeaturesEXT;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`depthBiasControl` indicates whether the implementation supports the `vkCmdSetDepthBias2EXT` command and the `VkDepthBiasRepresentationInfoEXT` structure.
- []()`leastRepresentableValueForceUnormRepresentation` indicates whether the implementation supports using the `VK_DEPTH_BIAS_REPRESENTATION_LEAST_REPRESENTABLE_VALUE_FORCE_UNORM_EXT` depth bias representation.
- []()`floatRepresentation` indicates whether the implementation supports using the `VK_DEPTH_BIAS_REPRESENTATION_FLOAT_EXT` depth bias representation.
- []()`depthBiasExact` indicates whether the implementation supports forcing depth bias to not be scaled to ensure a minimum resolvable difference using `VkDepthBiasRepresentationInfoEXT`::`depthBiasExact`.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceDepthBiasControlFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceDepthBiasControlFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_BIAS_CONTROL_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_depth\_bias\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_bias_control.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceDepthBiasControlFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0