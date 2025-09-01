# VK\_SEC\_amigo\_profiling(3) Manual Page

## Name

VK\_SEC\_amigo\_profiling - device extension



## [](#_registered_extension_number)Registered Extension Number

486

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Ralph Potter \[GitLab]r\_potter

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-07-29

**IP Status**

No known IP claims.

**Contributors**

- Ralph Potter, Samsung
- Sangrak Oh, Samsung
- Jinku Kang, Samsung

## [](#_description)Description

This extension is intended to communicate information from layered API implementations such as ANGLE to internal proprietary system schedulers. It has no behavioral implications beyond enabling more intelligent behavior from the system scheduler.

Application developers should avoid using this extension. It is documented solely for the benefit of tools and layer developers, who may need to manipulate `pNext` chains that include these structures.

Note

There is currently no specification language written for this extension. The links to APIs defined by the extension are to stubs that only include generated content such as API declarations and implicit valid usage statements.

Note

This extension is only intended for use in specific embedded environments with known implementation details, and is therefore undocumented.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceAmigoProfilingFeaturesSEC](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceAmigoProfilingFeaturesSEC.html)
- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html):
  
  - [VkAmigoProfilingSubmitInfoSEC](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAmigoProfilingSubmitInfoSEC.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_SEC_AMIGO_PROFILING_EXTENSION_NAME`
- `VK_SEC_AMIGO_PROFILING_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_AMIGO_PROFILING_SUBMIT_INFO_SEC`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_AMIGO_PROFILING_FEATURES_SEC`

## [](#_version_history)Version History

- Revision 1, 2022-07-29 (Ralph Potter)
  
  - Initial specification

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_SEC_amigo_profiling).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0