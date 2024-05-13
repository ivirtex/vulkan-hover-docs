# VK_KHR_maintenance4(3) Manual Page

## Name

VK_KHR_maintenance4 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

414

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.3-promotions"
  target="_blank" rel="noopener">Vulkan 1.3</a>

## <a href="#_contact" class="anchor"></a>Contact

- Piers Daniell <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_maintenance4%5D%20@pdaniell-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_maintenance4%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pdaniell-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-10-25

**Interactions and External Dependencies**  
- Requires SPIR-V 1.2 for `LocalSizeId`

**Contributors**  
- Lionel Duc, NVIDIA

- Faith Ekstrand, Intel

- Spencer Fricke, Samsung

- Tobias Hector, AMD

- Lionel Landwerlin, Intel

- Graeme Leese, Broadcom

- Tom Olson, Arm

- Stu Smith, AMD

- Yiwei Zhang, Google

## <a href="#_description" class="anchor"></a>Description

`VK_KHR_maintenance4` adds a collection of minor features, none of which
would warrant an entire extension of their own.

The new features are as follows:

- Allow the application to destroy their
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) object immediately after it
  was used to create another object. It is no longer necessary to keep
  its handle valid while the created object is in use.

- Add a new <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxBufferSize"
  target="_blank" rel="noopener"><code>maxBufferSize</code></a>
  implementation-defined limit for the maximum size `VkBuffer` that
  **can** be created.

- Add support for the SPIR-V 1.2 `LocalSizeId` execution mode, which can
  be used as an alternative to `LocalSize` to specify the local
  workgroup size with specialization constants.

- Add a guarantee that images created with identical creation parameters
  will always have the same alignment requirements.

- Add new
  [vkGetDeviceBufferMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceBufferMemoryRequirementsKHR.html),
  [vkGetDeviceImageMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageMemoryRequirementsKHR.html),
  and
  [vkGetDeviceImageSparseMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageSparseMemoryRequirementsKHR.html)
  to allow the application to query the image memory requirements
  without having to create an image object and query it.

- Relax the requirement that push constants must be initialized before
  they are dynamically accessed.

- Relax the interface matching rules to allow a larger output vector to
  match with a smaller input vector, with additional values being
  discarded.

- Add a guarantee for buffer memory requirement that the size memory
  requirement is never greater than the result of aligning create size
  with the alignment memory requirement.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetDeviceBufferMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceBufferMemoryRequirementsKHR.html)

- [vkGetDeviceImageMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageMemoryRequirementsKHR.html)

- [vkGetDeviceImageSparseMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageSparseMemoryRequirementsKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkDeviceBufferMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceBufferMemoryRequirementsKHR.html)

- [VkDeviceImageMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceImageMemoryRequirementsKHR.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceMaintenance4FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance4FeaturesKHR.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceMaintenance4PropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance4PropertiesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_MAINTENANCE_4_EXTENSION_NAME`

- `VK_KHR_MAINTENANCE_4_SPEC_VERSION`

- Extending [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html):

  - `VK_IMAGE_ASPECT_NONE_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DEVICE_BUFFER_MEMORY_REQUIREMENTS_KHR`

  - `VK_STRUCTURE_TYPE_DEVICE_IMAGE_MEMORY_REQUIREMENTS_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_4_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_4_PROPERTIES_KHR`

## <a href="#_promotion_to_vulkan_1_3" class="anchor"></a>Promotion to Vulkan 1.3

Functionality in this extension is included in core Vulkan 1.3, with the
KHR suffix omitted. The original type, enum and command names are still
available as aliases of the core functionality.

## <a href="#_issues" class="anchor"></a>Issues

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-08-18 (Piers Daniell)

  - Internal revisions

- Revision 2, 2021-10-25 (Yiwei Zhang)

  - More guarantees on buffer memory requirements

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_maintenance4"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
