# VkConformanceVersion(3) Manual Page

## Name

VkConformanceVersion - Structure containing the conformance test suite version the implementation is compliant with



## [](#_c_specification)C Specification

The conformance test suite version an implementation is compliant with is described with the `VkConformanceVersion` structure:

```c++
// Provided by VK_VERSION_1_2
typedef struct VkConformanceVersion {
    uint8_t    major;
    uint8_t    minor;
    uint8_t    subminor;
    uint8_t    patch;
} VkConformanceVersion;
```

or the equivalent

```c++
// Provided by VK_KHR_driver_properties
typedef VkConformanceVersion VkConformanceVersionKHR;
```

## [](#_members)Members

- `major` is the major version number of the conformance test suite.
- `minor` is the minor version number of the conformance test suite.
- `subminor` is the subminor version number of the conformance test suite.
- `patch` is the patch version number of the conformance test suite.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_KHR\_driver\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_driver_properties.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkPhysicalDeviceDriverProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDriverProperties.html), [VkPhysicalDeviceVulkan12Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan12Properties.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkConformanceVersion).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0