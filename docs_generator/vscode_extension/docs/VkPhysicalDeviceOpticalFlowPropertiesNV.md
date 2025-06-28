# VkPhysicalDeviceOpticalFlowPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceOpticalFlowPropertiesNV - Structure describing properties supported by VK\_NV\_optical\_flow



## [](#_c_specification)C Specification

The `VkPhysicalDeviceOpticalFlowPropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_optical_flow
typedef struct VkPhysicalDeviceOpticalFlowPropertiesNV {
    VkStructureType                 sType;
    void*                           pNext;
    VkOpticalFlowGridSizeFlagsNV    supportedOutputGridSizes;
    VkOpticalFlowGridSizeFlagsNV    supportedHintGridSizes;
    VkBool32                        hintSupported;
    VkBool32                        costSupported;
    VkBool32                        bidirectionalFlowSupported;
    VkBool32                        globalFlowSupported;
    uint32_t                        minWidth;
    uint32_t                        minHeight;
    uint32_t                        maxWidth;
    uint32_t                        maxHeight;
    uint32_t                        maxNumRegionsOfInterest;
} VkPhysicalDeviceOpticalFlowPropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`supportedOutputGridSizes` are the supported [VkOpticalFlowGridSizeFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowGridSizeFlagsNV.html) which can be specified in `VkOpticalFlowSessionCreateInfoNV`::`outputGridSize`.
- []()`supportedHintGridSizes` are the supported [VkOpticalFlowGridSizeFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowGridSizeFlagsNV.html) which can be specified in `VkOpticalFlowSessionCreateInfoNV`::`hintGridSize`.
- []()`hintSupported` is a boolean describing whether using hint flow vector map is supported in an optical flow session.
- []()`costSupported` is a boolean describing whether cost map generation is supported in an optical flow session.
- []()`bidirectionalFlowSupported` is a boolean describing whether bi-directional flow generation is supported in an optical flow session.
- []()`globalFlowSupported` is a boolean describing whether global flow vector map generation is supported in an optical flow session.
- []()`minWidth` is the minimum width in pixels for images used in an optical flow session.
- []()`minHeight` is the minimum height in pixels for images used in an optical flow session.
- []()`maxWidth` is the maximum width in pixels for images used in an optical flow session.
- []()`maxHeight` is the maximum height in pixels for images used in an optical flow session.
- []()`maxNumRegionsOfInterest` is the maximum number of regions of interest which can be used in an optical flow session. If this `maxNumRegionsOfInterest` is 0, regions of interest are not supported in an optical flow session.

## [](#_description)Description

If the `VkPhysicalDeviceOpticalFlowPropertiesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceOpticalFlowPropertiesNV-sType-sType)VUID-VkPhysicalDeviceOpticalFlowPropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_OPTICAL_FLOW_PROPERTIES_NV`

## [](#_see_also)See Also

[VK\_NV\_optical\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_optical_flow.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkOpticalFlowGridSizeFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowGridSizeFlagsNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceOpticalFlowPropertiesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0