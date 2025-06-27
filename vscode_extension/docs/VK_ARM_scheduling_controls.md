# VK_ARM_scheduling_controls(3) Manual Page

## Name

VK_ARM_scheduling_controls - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

418

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_ARM_shader_core_builtins](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ARM_shader_core_builtins.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Kevin Petit <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_ARM_scheduling_controls%5D%20@kpet%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_ARM_scheduling_controls%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>kpet</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-08-23

**Interactions and External Dependencies**  
None

**IP Status**  
No known IP claims.

**Contributors**  
- Kévin Petit, Arm Ltd.

- Jan-Harald Fredriksen, Arm Ltd.

- Mikel Garai, Arm Ltd.

## <a href="#_description" class="anchor"></a>Description

This extension exposes a collection of controls to modify the scheduling
behavior of Arm Mali devices.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkDeviceQueueShaderCoreControlCreateInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueShaderCoreControlCreateInfoARM.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceSchedulingControlsFeaturesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSchedulingControlsFeaturesARM.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceSchedulingControlsPropertiesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSchedulingControlsPropertiesARM.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkPhysicalDeviceSchedulingControlsFlagBitsARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSchedulingControlsFlagBitsARM.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkPhysicalDeviceSchedulingControlsFlagsARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSchedulingControlsFlagsARM.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_ARM_SCHEDULING_CONTROLS_EXTENSION_NAME`

- `VK_ARM_SCHEDULING_CONTROLS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DEVICE_QUEUE_SHADER_CORE_CONTROL_CREATE_INFO_ARM`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SCHEDULING_CONTROLS_FEATURES_ARM`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SCHEDULING_CONTROLS_PROPERTIES_ARM`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

None.

## <a href="#_issues" class="anchor"></a>Issues

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-08-23 (Kévin Petit)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_ARM_scheduling_controls"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
