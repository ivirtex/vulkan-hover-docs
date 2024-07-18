# VK_ANDROID_external_format_resolve(3) Manual Page

## Name

VK_ANDROID_external_format_resolve - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

469

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_ANDROID_external_memory_android_hardware_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_VERSION_1_3

- Interacts with VK_KHR_dynamic_rendering

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">OpenGL / ES support</a>

## <a href="#_contact" class="anchor"></a>Contact

- Chris Forbes <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_ANDROID_external_format_resolve%5D%20@chrisforbes%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_ANDROID_external_format_resolve%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>chrisforbes</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_ANDROID_external_format_resolve](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_ANDROID_external_format_resolve.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension enables rendering to Android Hardware Buffers with
external formats which cannot be directly represented as renderable in
Vulkan, including Y′C<sub>B</sub>C<sub>R</sub> formats.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferPropertiesANDROID.html):

  - [VkAndroidHardwareBufferFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatResolvePropertiesANDROID.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceExternalFormatResolveFeaturesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalFormatResolveFeaturesANDROID.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceExternalFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalFormatResolvePropertiesANDROID.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_ANDROID_EXTERNAL_FORMAT_RESOLVE_EXTENSION_NAME`

- `VK_ANDROID_EXTERNAL_FORMAT_RESOLVE_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_FORMAT_RESOLVE_PROPERTIES_ANDROID`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FORMAT_RESOLVE_FEATURES_ANDROID`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FORMAT_RESOLVE_PROPERTIES_ANDROID`

If [VK_KHR_dynamic_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering.html) or [Version
1.3](#versions-1.3) is supported:

- Extending [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html):

  - `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-05-34 (Tobias Hector)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_ANDROID_external_format_resolve"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
