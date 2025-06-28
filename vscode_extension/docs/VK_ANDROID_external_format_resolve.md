# VK\_ANDROID\_external\_format\_resolve(3) Manual Page

## Name

VK\_ANDROID\_external\_format\_resolve - device extension



## [](#_registered_extension_number)Registered Extension Number

469

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_ANDROID\_external\_memory\_android\_hardware\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_3
- Interacts with VK\_KHR\_dynamic\_rendering

## [](#_special_use)Special Use

- [OpenGL / ES support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Chris Forbes [\[GitHub\]chrisforbes](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_ANDROID_external_format_resolve%5D%20%40chrisforbes%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_ANDROID_external_format_resolve%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_ANDROID\_external\_format\_resolve](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_ANDROID_external_format_resolve.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-05-03

**IP Status**

No known IP claims.

**Contributors**

- Tobias Hector, AMD
- Chris Forbes, Google
- Jan-Harald Fredriksen, Arm
- Shahbaz Youssefi, Google
- Matthew Netsch, Qualcomm
- Tony Zlatsinki, Nvidia
- Daniel Koch, Nvidia
- Jeff Leger, Qualcomm
- Alex Walters, Imagination
- Andrew Garrard, Imagination
- Ralph Potter, Samsung
- Ian Elliott, Google

## [](#_description)Description

This extension enables rendering to Android Hardware Buffers with external formats which cannot be directly represented as renderable in Vulkan, including Y′CBCR formats.

## [](#_new_structures)New Structures

- Extending [VkAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferPropertiesANDROID.html):
  
  - [VkAndroidHardwareBufferFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatResolvePropertiesANDROID.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceExternalFormatResolveFeaturesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalFormatResolveFeaturesANDROID.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceExternalFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalFormatResolvePropertiesANDROID.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_ANDROID_EXTERNAL_FORMAT_RESOLVE_EXTENSION_NAME`
- `VK_ANDROID_EXTERNAL_FORMAT_RESOLVE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_FORMAT_RESOLVE_PROPERTIES_ANDROID`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FORMAT_RESOLVE_FEATURES_ANDROID`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FORMAT_RESOLVE_PROPERTIES_ANDROID`

If [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html) or [Vulkan Version 1.3](#versions-1.3) is supported:

- Extending [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlagBits.html):
  
  - `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`
  - `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID`

## [](#_version_history)Version History

- Revision 1, 2023-05-34 (Tobias Hector)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_ANDROID_external_format_resolve)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0